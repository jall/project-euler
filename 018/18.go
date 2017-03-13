package main

import (
	"fmt"
)

// Thoughts:
//	 The data structure is not a binary tree, as nodes can have more than one parent
//   It is very similar to Pascal's Triangle

// Start from bottom+1 row, find max(self+left, self+right), replace self with result, continue along row,
// discard row below, move to row above, repeat.
// E.g. given:
//   A
//  B C
// D E F
// Compute:
// B = max(B+D,B+E)
// C = max(C+E,C+F)
// Discard DEF row
// A = max(A+B, A+C)
// Result = A = largest possible leaf to root sum
func main() {
	data := getData()

	for i := len(data) - 2; i >= 0; i-- {
		row := data[i]
		for j := range row {
			data[i][j] += max(data[i+1][j], data[i+1][j+1])
		}
	}

	fmt.Println(data[0][0])
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func getData() [][]int {
	return [][]int{
		{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 4, 82, 47, 65},
		{19, 1, 23, 75, 3, 34},
		{88, 2, 77, 73, 7, 63, 67},
		{99, 65, 4, 28, 6, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
	}
}
