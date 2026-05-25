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
func enesimo(n int, atual int, cont int) int {
	if eh_primo(atual, 2) {
		cont++
	}
	if cont == n {
		return atual
	}
	return enesimo(n, atual+1, cont)
}
func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(enesimo(x, 2, 0))
}
