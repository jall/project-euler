package numberutil

// TODO Tests

func ProperDivisors(number int) (divisors []int) {
	primeFactors := PrimeFactors(number)

	for divisor := range ProperDivisorsFromPrimeFactors(primeFactors) {
		divisors = append(divisors, divisor)
	}

	return divisors
}

func ProperDivisorsFromPrimeFactors(primeFactors []int) map[int]int {
	divisors := map[int]int{}

	for _, set := range PowerSetWithOptions(primeFactors, true, false) {
		divisor := 1
		for _, number := range set {
			divisor *= number
		}
		divisors[divisor]++
	}

	return divisors
}
