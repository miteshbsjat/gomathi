package gomathi

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
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
func Median[T Number](numbers []T) float64 {
	l := len(numbers)
	if l == 0 {
		return 0
	}
	// for data immutability
	nCopy := make([]T, l)
	copy(nCopy, numbers)
	slices.Sort(nCopy)
	var median float64
	var mid int = l / 2
	if l%2 == 0 {
		median = float64((nCopy[mid-1] + nCopy[mid]) / 2.0)
	} else {
		median = float64(nCopy[mid])
	}

	return median
}

// Find the pth-percentile value in numbers
func Quantile[T Number](p float32, numbers []T) T {
	l := len(numbers)
	if l == 0 {
		return 0
	}
	nCopy := make([]T, l)
	copy(nCopy, numbers)

	slices.Sort(nCopy)
	pIndex := int(p * float32(l))
	return nCopy[pIndex]
}
