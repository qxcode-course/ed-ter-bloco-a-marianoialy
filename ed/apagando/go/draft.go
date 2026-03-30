package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	fmt.Scan(&m)
	saiu := make([]int, m)

	for i := 0; i < m; i++ {
		fmt.Scan(&saiu[i])
	}
	ficou := make([]int, n)

	for i := 0; i < n; i++ {
		ficou[i] = vetor[i]
		for j := 0; j < m; j++ {
			if vetor[i] == saiu[j] {
				ficou[i] = 0
			}
		}
	}
	for i := 0; i < n; i++ {
		if (ficou[i]) == 0 {
			continue
		} else {
			fmt.Print(ficou[i], " ")
		}
	}
	fmt.Println()
}
