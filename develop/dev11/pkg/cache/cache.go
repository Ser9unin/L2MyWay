package calendar

import (
	"errors"
	"fmt"
	"time"

	model "github.com/Ser9unin/L2MyWay/develop/dev11/pkg/model"
)

// CacheCalendar - type of calendar cache
type CacheCalendar struct {
	users map[model.UserID]model.User
}

// NewCalendarCache - constructor
func NewUserCalendar() *CacheCalendar {
	return &CacheCalendar{
		users: *new(map[model.UserID]model.User),
	}
}

// CreateEvent - create event
func (cc *CacheCalendar) CreateEvent(userID model.UserID, date model.Dates, event model.Event) (*model.Event, error) {
	if _, ok := cc.users[userID]; !ok {
		newUser := model.User{}
		newUser.ID = userID
		newUser.EventsMap[date] = append(newUser.EventsMap[date], event)
		cc.users[userID] = newUser
	} else {
		cc.users[userID].EventsMap[date] = append(cc.users[userID].EventsMap[date], event)
	}

	return &event, nil
}

// UpdateEvent - update event
func (cc *CacheCalendar) UpdateEvent(userID model.UserID, date model.Dates, event model.Event) error {
	if _, ok := cc.users[userID]; !ok {
		errorText := fmt.Sprintf("user id %v not found", userID)
		return errors.New(errorText)
	} else {
		if _, ok := cc.users[userID].EventsMap[date]; !ok {
			errorText := fmt.Sprintf("events on date %v not found", date)
			return errors.New(errorText)
		}
		for i, val := range cc.users[userID].EventsMap[date] {
			if val.ID == event.ID {
				cc.users[userID].EventsMap[date][i] = event
			}
		}
	}

	return nil
}

// DeleteEvent - delete event
func (cc *CacheCalendar) DeleteEvent(userID model.UserID, date model.Dates, eventID model.EventID) error {
	if _, ok := cc.users[userID]; !ok {
		errorText := fmt.Sprintf("user id %v not found", userID)
		return errors.New(errorText)
	} else {
		if _, ok := cc.users[userID].EventsMap[date]; !ok {
			errorText := fmt.Sprintf("events on date %v not found", date)
			return errors.New(errorText)
		}
		for i, val := range cc.users[userID].EventsMap[date] {
			if val.ID == eventID {
				copy(cc.users[userID].EventsMap[date][i:], cc.users[userID].EventsMap[date][i+1:])
			}
		}
	}

	return nil
}

// GetEventByID - get event by id
func (cc *CacheCalendar) GetEventByID(userID model.UserID, eventID model.EventID) (*model.Event, error) {
	for _, datesval := range cc.users[userID].EventsMap {
		for _, eventsv := range datesval {
			if eventsv.ID == eventID {
				return &eventsv, nil
			}
		}
	}

	errorText := fmt.Sprintf("event with id %v not found", eventID)
	return nil, errors.New(errorText)
}

// GetEventsForDate - get all events for a given period of time
func (cc *CacheCalendar) GetEventsForDate(userID model.UserID, day time.Time) ([]model.Event, error) {
	var events []model.Event

	if _, ok := cc.users[userID]; !ok {
		errorText := fmt.Sprintf("user id %v not found", userID)
		return nil, errors.New(errorText)
	} else {
		for k, val := range cc.users[userID].EventsMap {
			dateToCheck := time.Time(k)
			if dateToCheck.Year() == day.Year() && dateToCheck.Month() == day.Month() && dateToCheck.Day() == day.Day() {
				events = append(events, val...)
			}
		}
	}
	return events, nil
}

func (cc *CacheCalendar) GetEventsForWeek(userID model.UserID, week time.Time) ([]model.Event, error) {
	var events []model.Event

	if _, ok := cc.users[userID]; !ok {
		errorText := fmt.Sprintf("user id %v not found", userID)
		return nil, errors.New(errorText)
	} else {
		for k, val := range cc.users[userID].EventsMap {
			dateToCheck := time.Time(k)
			cacheWeek, cacheYear := dateToCheck.ISOWeek()
			matchWeek, matchYear := week.ISOWeek()
			if cacheWeek == matchWeek && cacheYear == matchYear {
				events = append(events, val...)
			}
		}
	}
	return events, nil
}

func (cc *CacheCalendar) GetEventsForMonth(userID model.UserID, month time.Time) ([]model.Event, error) {
	var events []model.Event

	if _, ok := cc.users[userID]; !ok {
		errorText := fmt.Sprintf("user id %v not found", userID)
		return nil, errors.New(errorText)
	} else {
		for k, val := range cc.users[userID].EventsMap {
			dateToCheck := time.Time(k)
			cacheMonth := dateToCheck.Month()
			cacheYear := dateToCheck.Year()
			matchMonth := month.Month()
			matchYear := month.Year()
			if cacheMonth == matchMonth && cacheYear == matchYear {
				events = append(events, val...)
			}
		}
	}
	return events, nil
}
