package main

import (
	"bufio"
	"fmt"
	"os"
)

func valido(s []byte, pos int, num byte, L int) bool {
	for i := pos - 1; i >= 0 && pos-i <= L; i-- {
		if s[i] == num {
			return false
		}
	}
	for i := pos + 1; i < len(s) && i-pos <= L; i++ {
		if s[i] == num {
			return false
		}
	}
	return true
}
func backtrack(s []byte, pos int, L int) bool {

	if pos == len(s) {
		return true
	}

	if s[pos] != '.' {

		if !valido(s, pos, s[pos], L) {
			return false
		}

		return backtrack(s, pos+1, L)
	}

	for num := byte('0'); num <= byte('0'+L); num++ {

		if valido(s, pos, num, L) {
			s[pos] = num
			if backtrack(s, pos+1, L) {
				return true
			}
			s[pos] = '.'
		}
	}
	return false
}
func main() {

	in := bufio.NewReader(os.Stdin)
	var str string
	var L int

	fmt.Fscan(in, &str)
	fmt.Fscan(in, &L)
	s := []byte(str)

	backtrack(s, 0, L)
	fmt.Println(string(s))
}
