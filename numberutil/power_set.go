package numberutil

import (
	"github.com/jall/project-euler/stringutil"
	"strconv"
	"strings"
)

// TODO Tests for these

func PowerSet(set []int) [][]int {
	return PowerSetWithOptions(set, true, true)
}

// Lovely algorithm from Arturo Magidin on StackExchange
// http://math.stackexchange.com/a/92087/377473
func PowerSetWithOptions(set []int, includeEmptySet, includeOriginalSet bool) [][]int {
	var powerSet [][]int
	var start, end int
	cardinality := len(set)

	if includeEmptySet {
		start = 0
	} else {
		start = 1
	}

	if includeOriginalSet {
		end = Pow(2, cardinality)
	} else {
		end = Pow(2, cardinality) - 1
	}

	// Trivial subsets:
	// 	 Starting from 0 includes the empty set
	//   Ending at 2^k includes the original set
	// To avoid both, go from 1 => 2^k -1
	for i := start; i <= end; i++ {
		var subSet []int
		binary := stringutil.LeftPad(strconv.FormatInt(int64(i), 2), "0", cardinality)
		for key, rune := range strings.Split(binary, "") {
			bit, _ := strconv.Atoi(rune)
			if bit == 1 {
				subSet = append(subSet, set[key])
			}
		}
		powerSet = append(powerSet, subSet)
	}

	return powerSet
}
