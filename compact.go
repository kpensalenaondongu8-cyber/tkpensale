package piscine

func Compact(ptr *[]string) int {
	compactCount := 0
	compactSlice := []string{}
	for _, val := range *ptr {
		if val != "" {
			compactSlice = append(compactSlice, val)
			compactCount++
		}
	}
	*ptr = compactSlice

	return compactCount
}
