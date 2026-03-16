package main

import "fmt"

func main() {
	var h, p, f, d, posicao int
	fmt.Scan(&h)
	fmt.Scan(&p)
	fmt.Scan(&f)
	fmt.Scan(&d)
	posicao = f

	for i := 0; i < 16; i++ {
		posicao = posicao + d
		if posicao < 0 {
			posicao = 15
		}
		if posicao > 15 {
			posicao = 0
		}
		if posicao == h {
			fmt.Println("S")
			return
		}
		if posicao == p {
			fmt.Println("N")
			return
		}
	}
}
