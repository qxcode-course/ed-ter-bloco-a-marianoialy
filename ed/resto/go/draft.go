package main

import "fmt"

func processo(n int) {
	if n == 0 {
		return
	}
	resto := n % 2
	divisao := n / 2

	processo(divisao)

	fmt.Println(divisao, resto)
}

func main() {
	var n int
	fmt.Scan(&n)
	processo(n)
}
