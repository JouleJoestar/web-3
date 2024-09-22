package main

import (
	"fmt"
	"math"
)

func M(p float64, v float64) float64 {
	var m float64 = p * v

	return m
}

func W(k float64, p float64, v float64) float64 {
	var w float64 = math.Sqrt(k / M(p, v))

	return w
}

func T(k float64, p float64, v float64) float64 {
	var t float64 = 6 / W(k, p, v)

	return t
}

func main() {
	var result, num1, num2, num3 float64
	fmt.Scanf("%f %f %f", &num1, &num2, &num3)
	result = T(num1, num2, num3)
	fmt.Println(result)
}
