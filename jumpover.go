package piscine

func JumpOver(str string) string {
	thoresult := ""
	for i, char := range str {
		if (i+1)%3 == 0 {
			thoresult += string(char)
		}
	}
	if len(thoresult) == 0 {
		return "\n"
	}
	return thoresult + "\n"
}
