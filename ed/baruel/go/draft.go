package main

import "fmt"

func main() {
	var quantfig int
	var quantpossui int
	fmt.Scan(&quantfig)
	fmt.Scan(&quantpossui)
	vetor := make([]int, quantpossui)
	repetidas := make([]int, quantpossui)
	var quantrep int
	for i := 0; i < quantpossui; i++ {
		fmt.Scan(&vetor[i])
	}
	for i := 0; i < quantpossui; i++ {
		if vetor[i] == 0 {
			continue
		}
		for j := i + 1; j < quantpossui; j++ {
			if vetor[i] == vetor[j] {
				repetidas[quantrep] = vetor[j]
				quantrep++
				vetor[j] = 0
			}
		}
	}
	if quantrep == 0 {
		fmt.Println("N")
	} else {
		for i := 0; i < quantrep; i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(repetidas[i])
		}
		fmt.Println()
	}

	faltantes := make([]int, quantfig)
	quantfalt := 0
	for i := 1; i <= quantfig; i++ {
		count := 0
		for j := 0; j < quantpossui; j++ {
			if vetor[j] == i {
				count++
				break
			}
		}
		if count == 0 {
			faltantes[quantfalt] = i
			quantfalt++
		}
	}
	if quantfalt == 0 {
		fmt.Println("N")
	} else {
		for i := 0; i < quantfalt; i++ {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(faltantes[i])
		}
		fmt.Println()
	}
}
