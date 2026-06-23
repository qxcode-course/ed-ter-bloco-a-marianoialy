package main

import "fmt"

func diagonal(s string, i int) {
	if i == len(s) {
		return
	}

	for j := 0; j < i; j++ {
		fmt.Print(" ")
	}
	fmt.Println(string(s[i]))
	diagonal(s, i+1)

}

func main() {
	var s string
	fmt.Scan(&s)
	diagonal(s, 0)
}
