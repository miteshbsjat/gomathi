package gomathi

import (
	"sort"
)

// Statistics:
// Find the mean of the n numbers
func Mean(numbers ...float64) float64 {
	l := (float64)(len(numbers))
	if l == 0 {
		return 0
	}
	var sum float64 = 0
	for _, n := range numbers {
		sum += n
	}
	return sum / l
}

// Statistics:
// Find the median of the n numbers
func Median(numbers ...float64) float64 {
	l := len(numbers)
	if l == 0 {
		return 0
	}
	dataCopy := make([]float64, l)
	copy(dataCopy, numbers)

	sort.Float64s(dataCopy)

	var median float64
	if l%2 == 0 {
		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
	} else {
		median = dataCopy[l/2]
	}

	return median
}
