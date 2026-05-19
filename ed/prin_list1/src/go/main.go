package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	valores := []string{}

	for n := l.Front(); n != l.End(); n = n.Next() {
		if n == sword {
			valores = append(valores, ">"+fmt.Sprint(n.Value))
		} else {
			valores = append(valores, fmt.Sprint(n.Value))
		}
	}

	return "[" + strings.Join(valores, " ") + "]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	it = it.Next()

	if it == l.End() {
		it = l.Front()
	}

	return it
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)

	l := NewDList[int]()

	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}

	sword := l.Front()

	for range chosen - 1 {
		sword = Next(l, sword)
	}

	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))

		l.Erase(Next(l, sword))

		sword = Next(l, sword)
	}

	fmt.Println(ToStr(l, sword))
}
