package model

import (
	"errors"
	"net/http"
	"time"
)

type Dates time.Time
type UserID uint64
type EventID uint64

type User struct {
	ID        UserID            `json:"id,omitempty"`
	EventsMap map[Dates][]Event `json:"events,omitempty"`
}

type Event struct {
	ID          EventID `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description"`
	Date        Dates   `json:"date,omitempty"`
}

type EventRepository interface {
	Create(user_id uint64, e Event) (Event, error)
	Update(user_id uint64, e Event) error
	Delete(user_id uint64, event_id uint64) error
	GetForDay(user_id uint64, day time.Time) ([]Event, error)
	GetForWeek(user_id uint64, week time.Time) ([]Event, error)
	GetForMonth(user_id uint64, month time.Time) ([]Event, error)
}

var (
	ErrNotFound            = errors.New("your requested item is not found")
	ErrInternalServerError = errors.New("internal server error")
)

// GetStatusCode gets http code from error
func GetStatusCode(err error) int {
	if errors.Is(err, ErrNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
