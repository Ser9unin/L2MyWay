package API

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	calendar "github.com/Ser9unin/L2MyWay/develop/dev11/pkg/cache"
	"github.com/Ser9unin/L2MyWay/develop/dev11/pkg/middleware"
	model "github.com/Ser9unin/L2MyWay/develop/dev11/pkg/model"
	"github.com/Ser9unin/L2MyWay/develop/dev11/pkg/render"
	"go.uber.org/zap"
)

type API struct {
	eventStore calendar.CacheInterface
	logger     *zap.Logger
}

func NewAPI(cache calendar.CacheInterface, logger *zap.Logger) API {
	return API{
		eventStore: cache,
		logger:     logger,
	}
}

func (a *API) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/create_event", middleware.Logger(a.Create))
	mux.HandleFunc("/update_event", middleware.Logger(a.Update))
	mux.HandleFunc("/delete_event", middleware.Logger(a.Delete))
	mux.HandleFunc("/events_for_day", middleware.Logger(a.Get))
	mux.HandleFunc("/events_for_week", middleware.Logger(a.Get))
	mux.HandleFunc("/events_for_month", middleware.Logger(a.Get))

	return mux
}

func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be post")
		return
	}

	err := r.ParseForm()
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse form")
		return
	}

	uid := r.FormValue("user_id")
	userID, err := strconv.Atoi(uid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user_id")
		return
	}
	userIdModel := model.UserID(userID)

	date := r.FormValue("date")
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use RFC3339 format")
		return
	}
	dateModel := model.Dates(t)

	title := r.FormValue("title")
	if title == "" {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("empty title"), "no title provided")
		return
	}

	e := model.Event{
		Title: title,
		Date:  dateModel,
	}

	result, err := a.eventStore.CreateEvent(userIdModel, dateModel, e)
	if err != nil {
		render.ErrorJSON(w, r, model.GetStatusCode(err), err, "can't create event")
		return
	}

	render.JSON(w, r, http.StatusCreated, result)
}

func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be put")
		return
	}

	err := r.ParseForm()
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse form")
		return
	}

	uid := r.FormValue("user_id")
	userID, err := strconv.Atoi(uid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user_id")
		return
	}
	userIdModel := model.UserID(userID)

	eid := r.FormValue("id")
	eventID, err := strconv.Atoi(eid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse id")
		return
	}
	eventIdModel := model.EventID(eventID)

	date := r.FormValue("date")
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use RFC3339 format")
		return
	}
	dateModel := model.Dates(t)

	title := r.FormValue("title")
	if title == "" {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("empty title"), "no title provided")
		return
	}

	e := model.Event{
		ID:    eventIdModel,
		Title: title,
		Date:  dateModel,
	}

	err = a.eventStore.UpdateEvent(userIdModel, dateModel, e)
	if err != nil {
		render.ErrorJSON(w, r, model.GetStatusCode(err), err, "can't update event")
		return
	}

	render.NoContent(w, r)
}

func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be delete")
		return
	}

	err := r.ParseForm()
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse form")
		return
	}

	uid := r.FormValue("user_id")
	userID, err := strconv.Atoi(uid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user_id")
		return
	}
	userIdModel := model.UserID(userID)

	eid := r.FormValue("id")
	eventID, err := strconv.Atoi(eid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse id")
		return
	}
	eventIdModel := model.EventID(eventID)

	date := r.FormValue("date")
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use RFC3339 format")
		return
	}
	dateModel := model.Dates(t)

	err = a.eventStore.DeleteEvent(userIdModel, dateModel, eventIdModel)
	if err != nil {
		render.ErrorJSON(w, r, model.GetStatusCode(err), err, "can't delete event")
		return
	}

	render.NoContent(w, r)
}

func (a *API) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		render.ErrorJSON(w, r, http.StatusBadRequest, fmt.Errorf("bad method: %s", r.Method), "method should be get")
		return
	}

	uid := r.FormValue("user_id")
	userID, err := strconv.Atoi(uid)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse user_id")
		return
	}
	userIdModel := model.UserID(userID)

	date := r.FormValue("date")
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		render.ErrorJSON(w, r, http.StatusBadRequest, err, "can't parse date, use RFC3339 format")
		return
	}

	events := make([]model.Event, 0)
	switch r.URL.Path {
	case "/events_for_day":
		events, err = a.eventStore.GetEventsForDate(userIdModel, t)
	case "/events_for_week":
		events, err = a.eventStore.GetEventsForWeek(userIdModel, t)
	case "/events_for_month":
		events, err = a.eventStore.GetEventsForMonth(userIdModel, t)
	}

	if err != nil {
		render.ErrorJSON(w, r, model.GetStatusCode(err), err, "can't get events")
		return
	}

	if len(events) == 0 {
		render.NoContent(w, r)
		return
	}
	render.JSON(w, r, http.StatusOK, events)
}
