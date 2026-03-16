package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	menordif := 101
	indiceganhador := -1
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scan(&a)
		fmt.Scan(&b)

		if a >= 10 && b >= 10 {
			diferenca := a - b

			if diferenca < 0 {
				diferenca = diferenca * -1

			}
			if diferenca < menordif {
				menordif = diferenca
				indiceganhador = i
			}
		}
	}
	if indiceganhador == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(indiceganhador)
	}
}
