package main

import (
	"fmt"
	"os"
	"strconv"
	"vmas/advent2024/utils"
)

var grid [][]int

type cell struct{ row, coll int }

func (c cell) isOut() bool {
	return c.row < 0 || c.row > len(grid)-1 || c.coll < 0 || c.coll > len(grid[0])-1
}

func (c cell) Next() []cell {
	return Next(c.row, c.coll)
}

func (c cell) Val() int {
	return grid[c.row][c.coll]
}

func Next(row, col int) []cell {
	next := make([]cell, 0)

	up := cell{row - 1, col}
	down := cell{row + 1, col}
	left := cell{row, col - 1}
	right := cell{row, col + 1}

	if !up.isOut() {
		next = append(next, up)
	}

	if !down.isOut() {
		next = append(next, down)
	}

	if !left.isOut() {
		next = append(next, left)
	}

	if !right.isOut() {
		next = append(next, right)
	}

	return next
}

func main() {
	lines := utils.ReadLines(os.Args[1])
	grid = make([][]int, len(lines))
	for row, line := range lines {
		grid[row] = make([]int, len(line))
		for col, val := range line {
			intVal := utils.Must(strconv.Atoi(string(val)))
			grid[row][col] = intVal
		}
	}

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
	c := cell{row, col}
	visited := make(map[cell]bool)
	visited[c] = true
	found := 0
	q := []cell{c}

	for len(q) > 0 {
		item := q[0]
		q = q[1:]

		if item.Val() == 9 {
			found += 1
		}

		for _, next := range item.Next() {
			validNext := next.Val()-item.Val() == 1
			if _, seen := visited[next]; validNext && !seen {
				visited[item] = true
				q = append(q, next)
			}
		}
	}

	return found
}
