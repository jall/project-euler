package numberutil

func CollatzSequence(n int) []int {
	return CollatzSequenceRecursive(n, []int{})
}

func CollatzSequenceRecursive(n int, sequence []int) []int {
	sequence = append(sequence, n)

	if n == 1 {
		return sequence
	}

	if n%2 == 0 {
		return CollatzSequenceRecursive(n/2, sequence)
	} else {
		return CollatzSequenceRecursive(3*n+1, sequence)
	}
}
