package numberutil

// NB: This does not separate out the repeating portion of the fraction from the rest.
// E.g. 1/12 expands to 0.08(3), and the expansion will show [0,8,3], but it doesn't
// tell you that only the 3 is the repeating portion.
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
