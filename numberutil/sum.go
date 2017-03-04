package numberutil

func Sum(numbers []int) int {
	total := 0

	for _, number := range numbers {
		total = total + number
	}

	return total
}
