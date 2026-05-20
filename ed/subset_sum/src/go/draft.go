package main

import "fmt"

func solve(v []int, i int, soma int, alvo int) bool {
	if soma == alvo {
		return true
	}
	if i == len(v) {
		return false
	}
	if solve(v, i+1, soma+v[i], alvo) {
		return true
	}
	if solve(v, i+1, soma, alvo) {
		return true
	}
	return false
}

func main() {
	var n, alvo int
	fmt.Scan(&n)
	fmt.Scan(&alvo)
	v := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&v[i])
	}
	fmt.Println(solve(v, 0, 0, alvo))
}
