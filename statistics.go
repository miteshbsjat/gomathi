package gomathi

func mean(numbers ...float64) float64 {
	var sum float64 = 0
	var count float64 = 0
	for _, n := range numbers {
		sum += n
		count += 1
	}
	return sum / count
}
