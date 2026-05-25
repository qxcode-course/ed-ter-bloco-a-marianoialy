package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	freq := map[int]int{}
	for _, x := range vet {
		if x < 0 {
			x = -x
		}
		freq[x]++
	}
	var resp []Pair
	for stress, qtd := range freq {
		resp = append(resp, Pair{stress, qtd})
	}
	for i := 0; i < len(resp); i++ {
		for j := i + 1; j < len(resp); j++ {
			if resp[j].One < resp[i].One {
				resp[i], resp[j] = resp[j], resp[i]
			}
		}
	}
	return resp
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}
	var resp []Pair
	atual := vet[0]
	cont := 1
	for i := 1; i < len(vet); i++ {
		if vet[i] == atual {
			cont++
		} else {
			resp = append(resp, Pair{atual, cont})
			atual = vet[i]
			cont = 1
		}
	}
	resp = append(resp, Pair{atual, cont})
	return resp
}

func mnext(vet []int) []int {
	resp := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			esq := i > 0 && vet[i-1] < 0
			dir := i < len(vet)-1 && vet[i+1] < 0
			if esq || dir {
				resp[i] = 1
			}
		}
	}
	return resp
}

func alone(vet []int) []int {
	resp := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			esq := i > 0 && vet[i-1] < 0
			dir := i < len(vet)-1 && vet[i+1] < 0
			if !esq && !dir {
				resp[i] = 1
			}
		}
	}
	return resp
}

func couple(vet []int) int {

	homens := map[int]int{}
	mulheres := map[int]int{}
	for _, x := range vet {

		if x > 0 {
			homens[x]++
		} else {
			mulheres[-x]++
		}
	}
	casais := 0
	for stress := range homens {
		if homens[stress] < mulheres[stress] {
			casais += homens[stress]
		} else {
			casais += mulheres[stress]
		}
	}
	return casais
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}
	for i := 0; i < len(seq); i++ {

		if vet[pos+i] != seq[i] {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	for i := 0; i < len(vet); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	remove := map[int]bool{}

	for _, pos := range posList {
		remove[pos] = true
	}

	var resp []int
	for i := 0; i < len(vet); i++ {

		if !remove[i] {
			resp = append(resp, vet[i])
		}
	}

	return resp
}

func clear(vet []int, value int) []int {
	var resp []int
	for _, x := range vet {

		if x != value {
			resp = append(resp, x)
		}
	}
	return resp
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
