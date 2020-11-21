//Package number offers methods chaining
package number

import "math"

//Number is of type float64
type Number float64

//Operation represents a valid operation on number
type Operation func(num1 Number, num2 float64) Number

//Add returns a number that allows for further chaining
func (n Number) Add(num float64) (sum Number) {
	return Number(float64(n) + num)
}

//Substract returns a number that allows for further chaining
func (n Number) Substract(num float64) (sum Number) {
	return Number(float64(n) - num)
}

//Multiply returns a number that allows for further chaining
func (n Number) Multiply(num float64) (sum Number) {
	return Number(float64(n) * num)
}

//Divide returns a number that allows for further chaining
func (n Number) Divide(num float64) (sum Number) {
	return Number(float64(n) * 1 / num)
}

//Absolute returns the abs of said number
func (n Number) Absolute() (sum Number) {
	return Number(math.Abs(float64(n)))
}
