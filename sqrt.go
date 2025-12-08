package piscine

func Sqrt(nb int) int {
	if nb <= 0 {
		return 0
	}
	for i := 1; int64(i)*int64(i) <= int64(nb); i++ {
		if int64(i)*int64(i) == int64(nb) {
			return i
		}
	}
	return 0
}
