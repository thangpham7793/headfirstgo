package utils

//Sum sums all given elements
func Sum(nums ...float64) (sum float64) {
	for _, num := range nums {
		sum += num
	}
	return
}

//SumSlice sums all elements in a slice
func SumSlice(nums []float64) (sum float64) {
	for _, num := range nums {
		sum += num
	}
	return
}

//Average calculates the average of given elements
func Average(nums ...float64) (average float64) {
	return Sum(nums...) / float64(len(nums))
}

//AverageSlice calculates the average of nums in a slice
func AverageSlice(nums []float64) (average float64) {
	return Sum(nums...) / float64(len(nums))
}
