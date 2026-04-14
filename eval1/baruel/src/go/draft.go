package main

import "fmt"

func main() {
	var totalalbum int
	// var repetidas int
	fmt.Scan(&totalalbum)
	var quantpossui int
	fmt.Scan(&quantpossui)
	possui := make([]int, quantpossui)
	vetorep := make([]int, quantpossui)
	for i := 0; i < quantpossui; i++ {
		fmt.Scan(&possui[i])
	}
	for i := 0; i < quantpossui; i++ {
		for j := 0; j < totalalbum; j++ {
			if possui[i] == possui[j+1] {
				vetorep[i] = possui[i]
				possui[i] = vetorep[i]
			}
			break
		}
	}
	for i := 0; i < quantpossui; i++ {
		fmt.Print(vetorep[i])
	}
}
