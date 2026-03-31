package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	men := []int{}
	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			men = append(men, vet[i])
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {

	women := []int{}
	for i := 0; i < len(vet); i++ {
		if vet[i] < 0 && vet[i] > -10 {
			women = append(women, vet[i])
		}
	}
	return women
}

func sortVet(vet []int) []int {
	copia := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		copia[i] = vet[i]
	}

	for i := 0; i < len(copia)-1; i++ {
		for j := 0; j < len(copia)-i-1; j++ {
			if copia[j] > copia[j+1] {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}
	return copia

}

func sortStress(vet []int) []int {

	copia := make([]int, len(vet))
	for i := 0; i < len(vet); i++ {
		copia[i] = vet[i]
	}

	for i := 0; i < len(copia)-1; i++ {
		for j := 0; j < len(copia)-i-1; j++ {
			a := copia[j]
			b := copia[j+1]

			if a < 0 {
				a = -a
			}
			if b < 0 {
				b = -b
			}

			if a > b {
				copia[j], copia[j+1] = copia[j+1], copia[j]
			}
		}
	}

	return copia
}

func reverse(vet []int) []int {
	invertido := []int{}
	for i := len(vet) - 1; i >= 0; i-- {
		invertido = append(invertido, vet[i])
	}
	return invertido
}

func unique(vet []int) []int {
	unicos := []int{}
	for i := 0; i < len(vet); i++ {
		tem := false
		for j := 0; j < len(unicos); j++ {
			if vet[i] == unicos[j] {
				tem = true
				break
			}
		}
		if !tem {
			unicos = append(unicos, vet[i])
		}
	}
	return unicos
}

func repeated(vet []int) []int {
	repetido := map[int]bool{}
	result := []int{}
	for _, v := range vet {
		if repetido[v] {
			result = append(result, v)
		}
		repetido[v] = true
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
