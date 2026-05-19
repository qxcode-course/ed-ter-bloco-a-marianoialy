package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	lin, col int
}

func estaDentro(matriz [][]byte, p Pos) bool {
	return p.lin >= 0 &&
		p.lin < len(matriz) &&
		p.col >= 0 &&
		p.col < len(matriz[0])
}
func dfs(
	matriz [][]byte,
	p Pos,
	visitados map[Pos]bool,
) {
	if !estaDentro(matriz, p) ||
		matriz[p.lin][p.col] != '1' ||
		visitados[p] {
		return
	}
	visitados[p] = true
	dfs(matriz, Pos{p.lin + 1, p.col}, visitados)
	dfs(matriz, Pos{p.lin - 1, p.col}, visitados)
	dfs(matriz, Pos{p.lin, p.col + 1}, visitados)
	dfs(matriz, Pos{p.lin, p.col - 1}, visitados)
}

// Não modifique a assinatura da função numIslands
func numIslands(grid [][]byte) int {

	visitados := make(map[Pos]bool)

	cont := 0

	for l := 0; l < len(grid); l++ {

		for c := 0; c < len(grid[0]); c++ {

			p := Pos{l, c}

			if grid[l][c] == '1' &&
				!visitados[p] {

				dfs(grid, p, visitados)

				cont++
			}
		}
	}

	return cont
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
