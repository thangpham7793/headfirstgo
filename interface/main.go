package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/calendar"
)

func main() {
	// d := calendar.Date{1992, 6, 6}
	// err := d.SetYear(1993)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = d.SetMonth(7)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = d.SetDay(7)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(&d)
	d, err := calendar.NewDate(1993, 7, 7)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)
	fmt.Println(d.Day())
	fmt.Println(d.Month())
	fmt.Println(d.Year())
}
