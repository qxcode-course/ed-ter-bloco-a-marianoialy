package main

import "fmt"

func ponto(gas, dist []int) int {
	total := 0
	atual := 0
	start := 0

	for i := 0; i < len(gas); i++ {
		dif := gas[i] - dist[i]
		total += dif
		atual += dif
		if atual < 0 {
			start = i + 1
			atual = 0
		}
	}
	return start
}
func main() {
	var n int
	fmt.Scan(&n)
	gas := make([]int, n)
	dist := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&gas[i], &dist[i])
	}
	fmt.Println(ponto(gas, dist))
}
