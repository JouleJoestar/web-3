package main

import (
	"fmt"
)

func getMaxDigit(s string) int {
	maxDigit := 0
	for i := 0; i < len(s); i++ {
		digit := int(s[i] - '0')
		if digit > maxDigit {
			maxDigit = digit
		}
	}
	return maxDigit
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(getMaxDigit(s))
}
