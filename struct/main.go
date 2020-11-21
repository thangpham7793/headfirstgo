package main

import (
	"fmt"

	"github.com/thangpham7793/headfirstgo/struct/subscriber"
)

//o satisfies interface Stringer if it does what a stringer does
func humanPrint(o fmt.Stringer) {
	fmt.Printf("Here is the human version:\n%s", o.String())
}

func main() {
	d := subscriber.NewDefault("Jack", subscriber.Address{"25", "Main Str"})
	c := subscriber.New("Tom", 10.5, false, subscriber.Address{"15", "Main Str"})
	fmt.Printf("%#v\n", d.Address)
	fmt.Println(c.Address)
	humanPrint(d)
}
