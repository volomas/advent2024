package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

var grid [][]rune
var vals []cell

type cell struct{ row, col int }

func main() {
	lines := utils.ReadLines(os.Args[1])
	grid = utils.CreateGrid(len(lines), len(lines[0]), '.')
	vals = make([]cell, 0)

	for row, line := range lines {
		for col, val := range line {
			if val != '.' {
				vals = append(vals, cell{row, col})
				grid[row][col] = val
			}
		}
	}

	count := 0
	for i := 0; i < len(vals)-1; i++ {
		for j := i + 1; j < len(vals); j++ {
			v1 := ValueAt(vals[i])
			v2 := ValueAt(vals[j])
			if v1 == v2 {
				a1, a2 := Antinodes(vals[i], vals[j])

				if !IsOut(a1) {

					if ValueAt(a1) != '#' {
						count += 1
					}

					if ValueAt(a1) == '.' {
						grid[a1.row][a1.col] = '#'
					}
				}

				if !IsOut(a2) {

					if ValueAt(a2) != '#' {
						count += 1
					}

					if ValueAt(a2) == '.' {
						grid[a2.row][a2.col] = '#'
					}
				}

				fmt.Printf("Antinodes for %v and %v\n", vals[i], vals[j])
				utils.PrintGrid(grid)
				fmt.Println()
			}
		}
	}

	utils.PrintGrid(grid)
	fmt.Println(count)
}

func ValueAt(c cell) rune {
	return grid[c.row][c.col]
}

func IsOut(c cell) bool {
	row := c.row
	col := c.col
	return row < 0 || col < 0 || row > len(grid)-1 || col > len(grid[0])-1
}

func Antinodes(cell1, cell2 cell) (cell, cell) {
	rowDiff := utils.Difference(cell1.row, cell2.row)
	collDiff := utils.Difference(cell1.col, cell2.col)

	//todo ugly, i know
	if cell1.row <= cell2.row {
		nr := cell1.row - rowDiff
		if cell1.col <= cell2.col {
			nc := cell1.col - collDiff
			return cell{nr, nc},
				cell{cell2.row + rowDiff, cell2.col + collDiff}
		} else {
			nc := cell1.col + collDiff
			return cell{nr, nc},
				cell{cell2.row + rowDiff, cell2.col - collDiff}
		}
	} else {
		nr := cell1.row + rowDiff
		if cell1.col <= cell2.col {
			nc := cell1.col - collDiff
			return cell{nr, nc},
				cell{cell2.row - rowDiff, cell2.col + collDiff}
		} else {
			nc := cell1.col + collDiff
			return cell{nr, nc},
				cell{cell2.row - rowDiff, cell2.col - collDiff}
		}
	}
}
