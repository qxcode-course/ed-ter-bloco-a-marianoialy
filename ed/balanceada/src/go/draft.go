package main

import (
	"fmt"
)

func balanceado(s string) bool {
	pilha := []rune{}

	for _, c := range s {
		if c == '(' || c == '[' {
			pilha = append(pilha, c)
			continue
		}
		if len(pilha) == 0 {
			return false
		}
		topo := pilha[len(pilha)-1]
		pilha = pilha[:len(pilha)-1]
		if c == ')' && topo != '(' {
			return false
		}
		if c == ']' && topo != '[' {
			return false
		}
	}
	return len(pilha) == 0
}

func main() {
	var s string
	fmt.Scan(&s)
	if balanceado(s) {
		fmt.Println("balanceado")
	} else {
		fmt.Println("nao balanceado")
	}
}
