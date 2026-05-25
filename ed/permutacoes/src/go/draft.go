package main

import (
	"fmt"
	"sort"
)

func permutar(letras []rune, usada []bool, atual string) {
	if len(atual) == len(letras) {
		fmt.Println(atual)
		return
	}
	for i := 0; i < len(letras); i++ {
		if !usada[i] {
			usada[i] = true
			permutar(letras, usada, atual+string(letras[i]))
			usada[i] = false
		}
	}
}
func main() {
	var s string
	fmt.Scan(&s)

	letras := []rune(s)

	sort.Slice(letras, func(i, j int) bool {
		return letras[i] < letras[j]
	})

	usada := make([]bool, len(letras))

	permutar(letras, usada, "")
}
