package calendar

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

//Event defines API of an event instance
type Event struct {
	title string
	Date
}

//NewEvent inits a new event and also calls constructor of Date
func NewEvent(title string, year, month, day int) (*Event, error) {
	e := Event{}
	err := e.SetTitle(title)
	if err != nil {
		return &e, err
	}
	d, err := NewDate(year, month, day)
	if err != nil {
		return &e, err
	}
	e.Date = *d
	return &e, nil
}

func (e Event) String() string {
	return fmt.Sprintf("%s %s", e.title, e.Date.String())
}

//SetTitle sets event title
func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("title too long")
	}
	e.title = title
	return nil
}

//Title gets title
func (e *Event) Title() string {
	return e.title
}

//DaysLeft returns the number of days left until the event's anniversary
func (e *Event) DaysLeft() string {
	current := time.Now()
	event := time.Date(time.Now().Year(), time.Month(e.month), e.day, 0, 0, 0, 0, time.Now().Location())
	diff := event.Sub(current).Hours() / 24
	if diff < 0 {
		event = time.Date(time.Now().Year()+1, time.Month(e.month), e.day, 0, 0, 0, 0, time.Now().Location())
		diff = event.Sub(current).Hours() / 24
	}
	return fmt.Sprintf("%.0f", diff)
}
