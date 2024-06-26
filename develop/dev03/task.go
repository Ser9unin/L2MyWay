package main

import (
	"log"
	"os"

	mysort "github.com/Ser9unin/L2MyWay/develop/dev03/pkg"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// собираем конфиг для сортировки
	args := os.Args[1:]
	var config mysort.SortCfg

	filepath, err := config.ParseConfig(args)
	if err != nil {
		log.Fatal(err)
	}

	// Читаем строки из файла
	data, err := mysort.GetDatafromFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	mysort.MySort(config, data)
}
