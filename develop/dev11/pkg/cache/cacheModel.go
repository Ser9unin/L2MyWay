package calendar

import (
	"time"

	model "github.com/Ser9unin/L2MyWay/develop/dev11/pkg/model"
)

// CacheInterface - interface for cache
type CacheInterface interface {
	CreateEvent(userID model.UserID, date model.Dates, event model.Event) (*model.Event, error)
	UpdateEvent(userID model.UserID, date model.Dates, event model.Event) error
	DeleteEvent(userID model.UserID, date model.Dates, eventID model.EventID) error
	GetEventByID(userID model.UserID, eventID model.EventID) (*model.Event, error)

	GetEventsForDate(userID model.UserID, date time.Time) ([]model.Event, error)
	GetEventsForWeek(userID model.UserID, date time.Time) ([]model.Event, error)
	GetEventsForMonth(userID model.UserID, date time.Time) ([]model.Event, error)
}

// Cache - base type of cache
type Cache struct {
	CacheInterface
}

// NewCache - constructor
func NewCache() *Cache {
	return &Cache{
		CacheInterface: NewUserCalendar(),
	}
}
