package utils

//SliceIndex returns ele index if found else -1
func SliceIndex(slice []string, ele string) int {
	for i, v := range slice {
		if v == ele {
			return i
		}
	}
	return -1
}
