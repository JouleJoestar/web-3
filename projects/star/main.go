package main

import (
	"fmt"
	"strings"
)

func UpdateString(str string) string {
	var newStr strings.Builder
	for i := 0; i < len(str); i++ {
		var letter string = string(str[i])
		newStr.WriteString(letter)
		if i < len(str)-1 {
			newStr.WriteString("*")
		}
	}
	return newStr.String()
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(UpdateString(str))
}
