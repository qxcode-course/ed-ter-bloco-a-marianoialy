package main

import (
	"bufio"
	"fmt"
	"os"
)

func dfs(grid [][]byte, word string, i int, j int, idx int) bool {
	if idx == len(word) {
		return true
	}
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}
	if grid[i][j] != word[idx] {
		return false
	}
	temp := grid[i][j]
	grid[i][j] = '#'
	found :=
		dfs(grid, word, i+1, j, idx+1) ||
			dfs(grid, word, i-1, j, idx+1) ||
			dfs(grid, word, i, j+1, idx+1) ||
			dfs(grid, word, i, j-1, idx+1)
	grid[i][j] = temp
	return found

}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if dfs(grid, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
