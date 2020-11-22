package main

import (
	"fmt"

	calc "github.com/headfirstgo/calcl"
	"github.com/headfirstgo/dates"
	"github.com/headfirstgo/greeting"
	"github.com/headfirstgo/greeting/vietnamese"
)

func main() {
	greeting.Hello()
	greeting.Hi()
	vietnamese.XinChao()
	fmt.Println(calc.Add(3.0, 4))
	fmt.Println(calc.Substract(3.0, 4))
	fmt.Println(calc.AddThenSubstract(7, 12.023))
	fmt.Printf("%d weeks are %d days\n", 2, dates.WeeksToDays(2))
}
