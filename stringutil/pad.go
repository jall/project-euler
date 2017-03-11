package stringutil

import "unicode/utf8"

func LeftPad(str string, pad string, length int) string {
	for {
		if utf8.RuneCountInString(str) >= length {
			return str
		}
		str = pad + str
	}
}
