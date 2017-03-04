package numberutil

func PrimeFactors(number int) []int {
	factors := make([]int, 0)
	divisor := 2

	// Pull out factors until we reach 1.
	for number > 1 {
		// Factor out divisor as many times as possible
		for number%divisor == 0 {
			factors = append(factors, divisor)
			number = number / divisor
		}
		divisor = divisor + 1
	}

	return factors
}

func PrimeFactorsGrouped(number int) map[int]int {
	factors := PrimeFactors(number)
	grouped := make(map[int]int)

	for _, factor := range factors {
		grouped[factor]++
	}

	return grouped
}
