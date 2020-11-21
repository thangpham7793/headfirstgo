package calc

//Add add two numbers
func Add(a, b float64) (sum float64) {
	return a + b
}

//Substract one num by another
func Substract(a, b float64) (remainder float64) {
	return a - b
}

//AddThenSubstract returns the first num
func AddThenSubstract(a, b float64) (original float64) {
	return (Add(a, b) + Substract(a, b)) / 2
}
