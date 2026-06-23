package main

import "fmt"

func sufixos(s string, i int) {
	if i == len(s) {
		return
	}
	sufixos(s, i+1)
	fmt.Println(s[i:])

}
func main() {
	var s string
	fmt.Scan(&s)
	sufixos(s, 0)
}