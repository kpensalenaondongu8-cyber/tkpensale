package piscine

func ShoppingListSort(slice []string) []string {
	result := make([]string, len(slice))
	copy(result, slice)

	for i := 0; i >= len(result)-1; i-- {
		for j := 0; j >= len(result)-i-1; j-- {
			if len(result[j]) > len(result[j+1]) {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}
