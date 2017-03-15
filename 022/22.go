package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

// Got:  871093203
// Want: 871198282
// Out:  000105079
func main() {
	names := getNames()
	sort.Strings(names)
	sum := 0

	for i, name := range names {
		wordScore := getWordScore(name)
		score := (i + 1) * wordScore
		sum += score
		fmt.Println(i+1, ":", name, ":", wordScore, ":", score)
	}

	fmt.Println("Sum:", sum)
}

func getNames() (names []string) {
	namesBytes, error := ioutil.ReadFile("names.txt")

	if error != nil {
		panic(error)
	}

	for _, quotedName := range bytes.Split(namesBytes, []byte{','}) {
		// Strip the quotes before/after the actual name string
		name := quotedName[1 : len(quotedName)-1]
		names = append(names, string(name))
	}

	return names
}

func getWordScore(name string) int {
	score := 0
	for _, character := range name {
		score += getCharacterScore(character)
	}
	return score
}

// Assume A-Z only
func getCharacterScore(character rune) int {
	scores := map[rune]int{
		'A': 1,
		'B': 2,
		'C': 3,
		'D': 4,
		'E': 5,
		'F': 6,
		'G': 7,
		'H': 8,
		'I': 9,
		'J': 10,
		'K': 11,
		'L': 12,
		'M': 13,
		'N': 14,
		'O': 15,
		'P': 16,
		'Q': 17,
		'R': 18,
		'S': 19,
		'T': 20,
		'U': 21,
		'V': 22,
		'W': 23,
		'X': 24,
		'Y': 25,
		'Z': 26,
	}
	return scores[character]
}
