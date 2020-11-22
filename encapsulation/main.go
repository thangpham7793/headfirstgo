package main

import (
	"fmt"

	"github.com/headfirstgo/calendar"
	"github.com/headfirstgo/general"
)

func main() {
	// d := calendar.Date{1992, 6, 6}
	// err := d.SetYear(1993)
	// general.HandleErr(err)
	// err = d.SetMonth(7)
	// general.HandleErr(err)
	// err = d.SetDay(7)
	// general.HandleErr(err)
	// fmt.Println(&d)
	// d, err := calendar.NewDate(1993, 7, 7)
	// general.HandleErr(err)
	// fmt.Println(d)
	// fmt.Println(d.Day())
	// fmt.Println(d.Month())
	// fmt.Println(d.Year())

	//these are just promoted methods from embedded type Date
	// e := calendar.Event{}
	// e.Title = "birthday"
	// e.SetYear(1993)
	// e.SetMonth(7)
	// e.SetDay(7)
	// fmt.Println(e)
	e, err := calendar.NewEvent("Birthday", 1993, 07, 07)
	general.HandleErr(err)
	fmt.Println(e)
	// fmt.Println(e.Title())
	// fmt.Println(e.Year())
	// fmt.Println(e.Month())
	// fmt.Println(e.Day())
	fmt.Println(e.DaysLeft())
	e, err = calendar.NewEvent("Upcoming Event", 1993, 12, 24)
	general.HandleErr(err)
	fmt.Println(e.DaysLeft())

}
