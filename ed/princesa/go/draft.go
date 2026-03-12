package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	var c float64
	var p float64
	var result float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	p = (a + b + c) / 2
	result = math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Printf("%.2f\n", result)
}
