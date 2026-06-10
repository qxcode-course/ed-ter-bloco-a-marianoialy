package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(matriz [][]byte, i, j int) {
	if i < 0 || i >= len(matriz) || j < 0 || j >= len(matriz[0]) {
		return
	}
	if matriz[i][j] != 'X' {
		return
	}
	matriz[i][j] = '.'
	dfs(matriz, i+1, j)
	dfs(matriz, i-1, j)
	dfs(matriz, i, j+1)
	dfs(matriz, i, j-1)
}

// Função que será chamada no LeetCode
func countBattleships(board [][]byte) int {
	cont := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				cont++
				dfs(board, i, j)
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

	board := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}

	result := countBattleships(board)
	fmt.Println(result)
}
