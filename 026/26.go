package main

import (
	"fmt"
	"github.com/jall/project-euler-go/numberutil"
)

func main() {
	longestRecurringCycle := []int{3}
	longestRecurringCycleDenominator := 3
	longestRecurringCycleLength := 1

	numerator := 1
	for denominator := 3; denominator < 1000; denominator++ {
		expansion, isRepeating := numberutil.ExpandFractionToDecimal(numerator, denominator)

		if isRepeating && len(expansion) > longestRecurringCycleLength {
			longestRecurringCycle = expansion
			longestRecurringCycleDenominator = denominator
			longestRecurringCycleLength = len(expansion)
		}
	}

	fmt.Println("Denominator:", longestRecurringCycleDenominator)
	fmt.Println("Cycle:", numberutil.IntegersToString(longestRecurringCycle))
	fmt.Println("Length:", longestRecurringCycleLength)
}
