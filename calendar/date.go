package calendar

import (
	"errors"
	"fmt"
)

//fields are unexported so cannot use literal syntax

//Date defines API of a date
type Date struct {
	year  int
	month int
	day   int
}

//NewDate inits a date
func NewDate(year, month, day int) (d *Date, err error) {

	d = &Date{1, 1, 1}

	err = d.SetYear(year)
	if err != nil {
		return nil, err
	}
	err = d.SetMonth(month)
	if err != nil {
		return nil, err
	}
	err = d.SetDay(day)
	if err != nil {
		return nil, err
	}
	return d, nil
}

//Stringer formats and returns the stored date
func (d *Date) String() string {
	return fmt.Sprintf("%d - %d - %d", d.day, d.month, d.year)
}

//SetYear sets the year
func (d *Date) SetYear(y int) error {
	switch {
	case y <= 0:
		return errors.New("year is negative")
	default:
		d.year = y
		return nil
	}
}

//SetMonth sets the month
func (d *Date) SetMonth(m int) error {
	switch {
	case m <= 0 || m > 12:
		return errors.New("invalid month")
	default:
		d.month = m
		return nil
	}
}

//SetDay sets the day
func (d *Date) SetDay(day int) error {
	switch {
	case day <= 0 || day >= 31:
		return errors.New("invalid day")
	default:
		d.day = day
		return nil
	}
}

//convention for getters is just the name of the field

//Year gets year
func (d *Date) Year() int {
	return d.year
}

//Month gets month
func (d *Date) Month() int {
	return d.month
}

//Day gets day
func (d *Date) Day() int {
	return d.day
}
