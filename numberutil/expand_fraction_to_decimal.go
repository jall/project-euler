package numberutil

func ExpandFractionToDecimal(numerator, denominator int) (expansion []int, isRepeating bool) {
	seenBefore := seenBeforeFactory()
	remainder := numerator % denominator

	for {
		// Terminating expansion
		if remainder == 0 {
			return expansion, false
		}

		// Repeating expansion
		if seenBefore(remainder, denominator) {
			return expansion, true
		}

		remainder *= 10
		expansion = append(expansion, remainder/denominator)
		remainder = remainder % denominator
	}
}

type RemainderSet struct {
	remainder   int
	denominator int
}

func seenBeforeFactory() func(remainder, denominator int) bool {
	previousRemainders := map[RemainderSet]bool{}

	return func(remainder, denominator int) bool {
		if _, isRepeating := previousRemainders[RemainderSet{remainder, denominator}]; isRepeating {
			return true
		} else {
			previousRemainders[RemainderSet{remainder, denominator}] = true
			return false
		}
	}
}
