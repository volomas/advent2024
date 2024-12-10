package main

import (
	"fmt"
	"os"
	u "vmas/advent2024/utils"
)

var grid u.Grid[int]

func main() {
	grid = u.ParseIntGrid(u.ReadLines(os.Args[1]))
	score := 0
	for row, line := range grid {
		for col, val := range line {
			if val == 0 {
				score += Score(row, col)
			}
		}
	}
	fmt.Println(score)
}

func Score(row, col int) int {
	c := u.NewCell(row, col)
	visited := make(map[u.Cell]bool)
	visited[c] = true
	found := 0
	q := []u.Cell{c}

	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		if grid.ValAt(item) == 9 {
			found += 1
		}

		for _, next := range grid.Next4(item) {
			validNext := grid.ValAt(next)-grid.ValAt(item) == 1
			if _, seen := visited[next]; validNext && !seen {
				visited[item] = true
				q = append(q, next)
			}
		}
	}

	return found
}
