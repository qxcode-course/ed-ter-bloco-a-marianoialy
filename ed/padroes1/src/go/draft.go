package main

import "fmt"

func bloquinhos(n int) int {
	if n == 1 {
		return 20
	}
	return bloquinhos(n-1) + 8
}
func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(bloquinhos(n))
}
