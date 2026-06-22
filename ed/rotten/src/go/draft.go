package main

import "fmt"

type Pos struct {
	l, c int
}

func (p Pos) getNeig() []Pos {
	return []Pos{
		{p.l + 1, p.c},
		{p.l - 1, p.c},
		{p.l, p.c + 1},
		{p.l, p.c - 1},
	}
}
func orangesRotting(grid [][]int) int {
	q := []Pos{}
	min := 0
	fresh := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 2 {
				q = append(q, Pos{i, j})
			}
			if grid[i][j] == 1 {
				fresh++
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	for len(q) > 0 {
		size := len(q)
		alterado := false
		for i := 0; i < size; i++ {
			atual := q[0]
			q = q[1:]

			for _, v := range atual.getNeig() {
				if v.l >= 0 && v.l < len(grid) && v.c >= 0 && v.c < len(grid[0]) && grid[v.l][v.c] == 1 {
					grid[v.l][v.c] = 2
					fresh--
					q = append(q, v)
					alterado = true
				}
			}
		}
		min++
		if !alterado {
			min--
			break
		}
	}
	if fresh > 0 {
		return -1
	}
	return min
}
func main() {
	var m, n int
	fmt.Scan(&m, &n)

	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	fmt.Println(orangesRotting(grid))
}
