package gomathi

// Find the mean of the n numbers
func Mean(numbers ...float64) float64 {
	var sum float64 = 0
	for _, n := range numbers {
		sum += n
	}
	return sum / (float64)(len(numbers))
}
