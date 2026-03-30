package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var e int
	fmt.Scan(&e)
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		vetor[i] = i + 1
	}
}
