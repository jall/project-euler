package numberutil

import (
	"bytes"
	"strconv"
)

func IntegersToString(slice []int) (str string) {
	var buffer bytes.Buffer

	for _, digit := range slice {
		buffer.WriteString(strconv.Itoa(digit))
	}

	return buffer.String()
}
