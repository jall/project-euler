package numberutil

func AliquotSum(number int) int {
	return Sum(ProperDivisors(number))
}

func IsPerfect(number int) bool {
	return number == AliquotSum(number)
}

func IsDeficient(number int) bool {
	return number > AliquotSum(number)
}

func IsAbundant(number int) bool {
	return number < AliquotSum(number)
}
