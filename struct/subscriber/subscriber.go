package subscriber

import "fmt"

//Subscriber contains fields for a subscriber
type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

//Employee contains fields for an employee
type Employee struct {
	Name   string
	Salary float64
	Address
}

//Address contains fields for an address
type Address struct {
	Number string
	Street string
}

//NewDefault creates a new sub with the given name
func NewDefault(name string, address Address) *Subscriber {
	return &Subscriber{name, 15.5, true, address}
}

//New creates custom sub
func New(name string, rate float64, active bool, address Address) *Subscriber {
	return &Subscriber{name, rate, active, address}
}

//String is called by fmt funcs //like toString() in other languages
func (s *Subscriber) String() string {
	return fmt.Sprintf("Name: %s\nRate: %.2f\nActive: %t\nAddress: %s %s\n", s.Name, s.Rate, s.Active, s.Number, s.Street)
}

//ApplyDiscount sets a lower rate for subscriber
func (s *Subscriber) ApplyDiscount(discount float64) {
	// (*s).Rate *= (1 - discount) no need to dereference since dot notation works on struct pointer as well
	s.Rate *= (1 - discount)
}
