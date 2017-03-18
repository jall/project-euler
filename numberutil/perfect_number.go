package numberutil

func IsPerfect(number int) bool {
	return number == Sum(ProperDivisors(number))
}
