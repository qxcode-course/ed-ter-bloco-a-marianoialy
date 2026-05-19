package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	linhas := len(grid)
	colunas := len(grid[0])

	var dfs func(int, int)

	dfs = func(l, c int) {
		if l < 0 || c < 0 || l >= linhas || c >= colunas {
			return
		}

		if grid[l][c] == '0' {
			return
		}

		grid[l][c] = '0'

		dfs(l+1, c)
		dfs(l-1, c)
		dfs(l, c+1)
		dfs(l, c-1)
	}

	qtd := 0

	for i := 0; i < linhas; i++ {
		for j := 0; j < colunas; j++ {
			if grid[i][j] == '1' {
				qtd++
				dfs(i, j)
			}
		}
	}

	return qtd
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
