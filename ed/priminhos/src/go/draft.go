package main

import "fmt"

func eh_primo(x int, div int) bool {
	if x < 2 {
		return false
	}
	if div*div > x {
		return true
	}
	if x%div == 0 {
		return false
	}
	return eh_primo(x, div+1)
}
func carregar_primos(n int, atual int, v []int) []int {
	if len(v) == n {
		return v
	}
	if eh_primo(atual, 2) {
		v = append(v, atual)
	}
	return carregar_primos(n, atual+1, v)
}
func main() {
	var n int
	fmt.Scan(&n)

	v := []int{}

	x := carregar_primos(n, 2, v)

	fmt.Print("[")

	for i := 0; i < len(x); i++ {
		if i > 0 {
			fmt.Print(", ")
		}

		fmt.Print(x[i])
	}

	fmt.Println("]")
}
