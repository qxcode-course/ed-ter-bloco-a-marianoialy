package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	vetor := make([]int, n)
	var result int
	for i := 0; i < n; i++ {
		fmt.Scan(&vetor[i])
	}
	for i := 0; i < n; i++ {
		if vetor[i] == 0 {
			continue
		}
		for j := 0; j < n; j++ {
			if vetor[j] == -vetor[i] {
				result++
				vetor[i] = 0
				vetor[j] = 0
				break
			}
		}

	}

	fmt.Println(result)
}
