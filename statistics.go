package gomathi

import (
	"sort"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Float | constraints.Integer
}

// Statistics:
// Find the mean of the n numbers
func Mean[T Number](numbers []T) float64 {
	l := (float64)(len(numbers))
	if l == 0 {
		return 0
	}
	var sum float64 = 0
	for _, n := range numbers {
		sum += float64(n)
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
	var mid int = l / 2
	if l%2 == 0 {
		median = (dataCopy[mid-1] + dataCopy[mid]) / 2
	} else {
		median = dataCopy[mid]
	}

	return median
}

// Find the pth-percentile value in numbers
func Quantile(p float32, numbers ...float64) float64 {
	l := len(numbers)
	if l == 0 {
		return 0
	}
	dataCopy := make([]float64, l)
	copy(dataCopy, numbers)

	sort.Float64s(dataCopy)
	pIndex := int(p * float32(l))
	sort.Float64s(dataCopy)
	return dataCopy[pIndex]
}
