package main

import "fmt"

func main() {
	var n string
	fmt.Scan(&n)
	for i := 0; i < len(n); i++ {
		var num int = int(n[i] - '0')
		var itog int = num * num
		fmt.Print(itog)
	}
	fmt.Println()
}
