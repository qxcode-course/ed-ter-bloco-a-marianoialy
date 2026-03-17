package main

import "fmt"

func main() {
	var quantfig int
	var quantpossui int
	fmt.Scan(&quantfig)
	fmt.Scan(&quantpossui)
	vetor := make([]int, quantpossui)
	var repetidas int
	var coladas int
	for i := 0; i < quantpossui; i++ {
		fmt.Scan(&vetor[i])
	}
	for i := 0; i < quantpossui; i++ {
		if vetor[i] == 0 {
			continue
		}
		for j := 0; j < quantpossui; j++ {
			if vetor[i] == vetor[j] {
				repetidas++
				vetor[j] = 0
			}
		}
		coladas++
	}
	fmt.Println(repetidas)
	fmt.Println(quantfig - coladas)
}
