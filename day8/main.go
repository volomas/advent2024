package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

var grid [][]rune
var vals []cell
var antinodes map[cell]bool

type cell struct{ row, col int }

func main() {
	lines := utils.ReadLines(os.Args[1])
	grid = utils.CreateGrid(len(lines), len(lines[0]), '.')
	vals = make([]cell, 0)
	antinodes = make(map[cell]bool)

	for row, line := range lines {
		for col, val := range line {
			if val != '.' {
				vals = append(vals, cell{row, col})
				grid[row][col] = val
			}
		}
	}

	for i := 0; i < len(vals)-1; i++ {
		for j := i + 1; j < len(vals); j++ {
			v1 := ValueAt(vals[i])
			v2 := ValueAt(vals[j])

			b1 := vals[i]
			b2 := vals[j]

			antinodes[b1] = true
			antinodes[b2] = true
			if v1 == v2 {
				a1, a2 := Antinodes(b1, b2)

				for !IsOut(a1) {
					if _, ok := antinodes[a1]; !ok {
						antinodes[a1] = true
					}

					if ValueAt(a1) == '.' {
						grid[a1.row][a1.col] = '#'
					}

					temp := a1
					a1, _ = Antinodes(a1, b1)
					b1 = temp
				}

				for !IsOut(a2) {
					if _, ok := antinodes[a2]; !ok {
						antinodes[a2] = true
					}

					if ValueAt(a2) == '.' {
						grid[a2.row][a2.col] = '#'
					}

					temp := a2
					_, a2 = Antinodes(b2, a2)
					b2 = temp
				}
			}
		}
	}

	utils.PrintGrid(grid)
	fmt.Println(len(antinodes))
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
