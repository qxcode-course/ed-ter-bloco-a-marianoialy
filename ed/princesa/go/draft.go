package main

import "fmt"

func main() {
	var n int
	var e int
	fmt.Scan(&n, &e)
	vetor := make([]int, n)
	for i := 0; i < n; i++ {
		vetor[i] = i + 1
	}
	pos := e - 1
	vivos := n

	for vivos > 0 {
		fmt.Print("[ ")
		for i := 0; i < n; i++ {
			if vetor[i] != 0 {
				if i == pos {
					fmt.Printf("%d> ", vetor[i])
				} else {
					fmt.Printf("%d ", vetor[i])
				}
			}
		}
		fmt.Println("]")
		if vivos == 1 {
			break
		}
		for {
			posMorto := (pos + 1) % n
			if vetor[posMorto] != 0 {
				vetor[posMorto] = 0
				vivos--
				break
			}
			pos++
		}
		for {
			pos = (pos + 1) % n
			if vetor[pos] != 0 {
				break
			}
		}
	}
}
