package numberutil

func MakeRange(min, max int) []int {
	slice := make([]int, max-min+1)
	for i := range slice {
		slice[i] = min + i
	}
	return slice
}
