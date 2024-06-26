package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/Ser9unin/L2MyWay/develop/dev11/cfg"
	api "github.com/Ser9unin/L2MyWay/develop/dev11/pkg/api"
	calendar "github.com/Ser9unin/L2MyWay/develop/dev11/pkg/cache"
	"go.uber.org/zap"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		os.Exit(1)
	}

	defer logger.Sync()

	zap.ReplaceGlobals(logger)

	logger.Info("reading config")

	config, err := config.NewConfig()
	if err != nil {
		logger.Error("can't decode config", zap.Error(err))
		return
	}

	store := calendar.NewCache()
	api := api.NewAPI(store, logger)
	router := api.NewRouter()

	srv := &http.Server{
		Addr:        config.HTTPServerAddress,
		Handler:     router,
		ReadTimeout: time.Duration(config.ReadTimeout) * time.Second,
		IdleTimeout: time.Duration(config.IdleTimeout) * time.Second,
	}

	logger.Info("running http server")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("can't start server", zap.Error(err), zap.String("server address", config.HTTPServerAddress))
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	logger.Info("received interrupt signal, closing server")
	timeout, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()
	if err := srv.Shutdown(timeout); err != nil {
		logger.Error("can't shutdown http server", zap.Error(err))
	}
}
