package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"regexp"
	"strconv"
	"strings"
)

// This whole question is a hacky piece of crap, but I'm not spending any more time redoing it!
func main() {
	characterCount := 0

	for i := 1; i <= 1000; i++ {
		str := convertToHumanReadable(i)
		regex, _ := regexp.Compile("[^A-Za-z]*")
		strippedString := regex.ReplaceAllString(str, "")
		characterCount += len(strippedString)
	}

	fmt.Println("Characters: ", characterCount)
}

func convertToHumanReadable(number int) string {
	if number == 0 {
		return "zero"
	}

	chunks := map[int]string{
		0: "",
		1: "",
		2: "",
		3: "",
	}

	digits := getDigits(number)

	for index, digit := range digits {
		isTeen := len(digits) > 1 && digits[1] == 1 && digits[0] != 0

		switch index {
		case 0:
			if !isTeen {
				chunks[0] = getUnitHumanReadable(digit)
			}
		case 1:
			if isTeen {
				chunks[1] = getTeenHumanReadable(digits[0])
			} else {
				chunks[1] = getTenHumanReadable(digit)
			}
		case 2:
			chunks[2] = getUnitHumanReadable(digit)
		case 3:
			chunks[3] = getUnitHumanReadable(digit)
		}
	}

	return buildLabelFromChunks(chunks)
}

func getDigits(number int) (digits []int) {
	characters := strings.Split(stringutil.Reverse(strconv.Itoa(number)), "")

	// Work through digits in reverse order
	for _, character := range characters {
		digit, _ := strconv.Atoi(character)
		digits = append(digits, digit)
	}

	return digits
}

func buildLabelFromChunks(chunks map[int]string) string {
	label := chunks[0]

	if label != "" {
		label = " " + label
	}

	label = chunks[1] + label

	if label != "" && chunks[2] != "" {
		label = " and " + label
	}

	if chunks[2] != "" {
		chunks[2] += " hundred"
	}

	label = chunks[2] + label

	if chunks[3] != "" && chunks[2] != "" {
		label = ", " + label
	}

	if chunks[3] != "" {
		chunks[3] += " thousand"
	}

	label = chunks[3] + label

	return label
}

func getUnitHumanReadable(number int) string {
	digits := map[int]string{
		0: "",
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	return digits[number]
}

func getTeenHumanReadable(number int) string {
	teens := map[int]string{
		1: "eleven",
		2: "twelve",
		3: "thirteen",
		4: "fourteen",
		5: "fifteen",
		6: "sixteen",
		7: "seventeen",
		8: "eighteen",
		9: "nineteen",
	}

	return teens[number]
}

func getTenHumanReadable(number int) string {
	digits := map[int]string{
		1: "ten",
		2: "twenty",
		3: "thirty",
		4: "fourty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninty",
	}

	return digits[number]
}
