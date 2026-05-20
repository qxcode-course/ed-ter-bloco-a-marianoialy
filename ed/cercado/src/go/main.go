package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	lin int
	col int
}

func dentro(board [][]byte, p Pos) bool {
	return p.lin >= 0 && p.lin < len(board) &&
		p.col >= 0 && p.col < len(board[0])
}

func dfs(board [][]byte, p Pos, visitados map[Pos]bool) {
	if !dentro(board, p) {
		return
	}

	if board[p.lin][p.col] != 'O' {
		return
	}

	if visitados[p] {
		return
	}

	visitados[p] = true

	dfs(board, Pos{p.lin - 1, p.col}, visitados)
	dfs(board, Pos{p.lin + 1, p.col}, visitados)
	dfs(board, Pos{p.lin, p.col - 1}, visitados)
	dfs(board, Pos{p.lin, p.col + 1}, visitados)
}

// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	nrows := len(board)
	ncols := len(board[0])

	visitados := make(map[Pos]bool)

	for col := 0; col < ncols; col++ {
		if board[0][col] == 'O' {
			dfs(board, Pos{0, col}, visitados)
		}

		if board[nrows-1][col] == 'O' {
			dfs(board, Pos{nrows - 1, col}, visitados)
		}
	}

	for lin := 0; lin < nrows; lin++ {
		if board[lin][0] == 'O' {
			dfs(board, Pos{lin, 0}, visitados)
		}

		if board[lin][ncols-1] == 'O' {
			dfs(board, Pos{lin, ncols - 1}, visitados)
		}
	}

	for lin := 0; lin < nrows; lin++ {
		for col := 0; col < ncols; col++ {
			p := Pos{lin, col}

			if board[lin][col] == 'O' && !visitados[p] {
				board[lin][col] = 'X'
			}
		}
	}
}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
