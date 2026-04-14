package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	var m int
	fmt.Scan(&m)
	deixou := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&deixou[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if vetor[i] == deixou[j] {
				vetor[i] = 0
				break
			}
		}
	}
	for i := 0; i < n; i++ {
		if vetor[i] != 0 {
			fmt.Print(vetor[i], " ")
		}

	}
	fmt.Println()
}
