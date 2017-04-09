package numberutil

func IsPrime(number int) bool {
	return len(PrimeFactors(number)) == 1
}
