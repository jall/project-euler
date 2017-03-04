package numberutil

// Integer power: compute a**b using binary powering algorithm
// See https://groups.google.com/forum/#!topic/golang-nuts/PnLnr4bc9Wo
func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}
