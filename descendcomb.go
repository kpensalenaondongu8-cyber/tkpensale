package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	first := true
	for a := 99; a >= 0; a-- {
		for b := a - 1; b >= 0; b-- {
			if !first {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
			first = false
			printTwoDigits(a)
			z01.PrintRune(' ')
			printTwoDigits(b)
		}
	}
}

func printTwoDigits(n int) {
	z01.PrintRune(rune(n/10 + '0'))
	z01.PrintRune(rune(n%10 + '0'))
}
