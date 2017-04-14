package numberutil

import (
	"strconv"
	"strings"
)

func Digits(number int) (digits []int) {
	for _, digitString := range strings.Split(strconv.Itoa(number), "") {
		digit, _ := strconv.Atoi(digitString)
		digits = append(digits, digit)
	}

	return digits
}
