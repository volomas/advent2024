package main

import (
	"fmt"
	"os"
	"slices"
	"vmas/advent2024/utils"
)

var garden utils.Grid[rune]

func main() {
	garden = utils.ParseCharGrid(utils.ReadLines(os.Args[1]))

	visited := make(map[utils.Cell]bool)
	result := 0
	for row, line := range garden {
		for col := range line {
			x := utils.Cell{Row: row, Coll: col}
			if _, ok := visited[x]; !ok {
				plot := FindPlot(visited, x)
				result += len(plot) * sides(plot)
			}
		}
	}

	fmt.Println(result)
}

// Find plot of the same plant (connected)
func FindPlot(visited map[utils.Cell]bool, from utils.Cell) []utils.Cell {
	plot := make([]utils.Cell, 0)
	visited[from] = true
	q := make([]utils.Cell, 0)
	q = append(q, from)
	for len(q) > 0 {
		item := q[0]
		q = q[1:]
		plot = append(plot, item)

		for _, n := range garden.Next4(item) {
			if garden.ValAt(n) != garden.ValAt(item) {
				continue
			}
			if _, ok := visited[n]; !ok {
				visited[n] = true
				q = append(q, n)
			}
		}
	}

	return plot
}

type Fence struct {
	utils.Cell
	side string
}

// calculate how many sides given plot has
func sides(plot []utils.Cell) int {
	plant := garden.ValAt(plot[0])
	fence := make([]Fence, 0)
	for _, c := range plot {

		up := Fence{utils.Cell{Row: c.Row - 1, Coll: c.Coll}, "up"}
		down := Fence{utils.Cell{Row: c.Row + 1, Coll: c.Coll}, "down"}
		left := Fence{utils.Cell{Row: c.Row, Coll: c.Coll - 1}, "left"}
		right := Fence{utils.Cell{Row: c.Row, Coll: c.Coll + 1}, "right"}

		for _, n := range []Fence{up, down, left, right} {
			if garden.IsOut(n.Cell) || garden.ValAt(n.Cell) != plant {
				fence = append(fence, n)
			}
		}
	}

	fenceBySide := make(map[string][]utils.Cell)
	for _, f := range fence {
		if fs, ok := fenceBySide[f.side]; ok {
			fenceBySide[f.side] = append(fs, f.Cell)

		} else {
			fenceBySide[f.side] = []utils.Cell{f.Cell}
		}

	}

	// fmt.Printf("Plant %c\n", plant)

	sidesCount := 0
	for _, side := range []string{"up", "down", "left", "right"} {

		//for up/down we nned to sort all cells by row, so that they apper one after another
		// [0, 1] -> [0, 2] -x [0, 4] (not touching)
		if side == "up" || side == "down" {
			slices.SortFunc(fence, compareByRow)
		} else {
			// [0,0] -> [1, 0] -x [3, 0] (not touching)
			slices.SortFunc(fence, compareByCol)
		}

		var prev *utils.Cell
		for _, f := range fence {
			if f.side != side {
				continue
			}

			if prev == nil || !touching(*prev, f.Cell) {
				sidesCount += 1
			}

			prev = &f.Cell
		}
	}
	return sidesCount

}

func touching(prevCell utils.Cell, thisCell utils.Cell) bool {
	rowDiff := utils.Difference(prevCell.Row, thisCell.Row)
	colDiff := utils.Difference(prevCell.Coll, thisCell.Coll)
	touching := (rowDiff == 1 && colDiff == 0) || (rowDiff == 0 && colDiff == 1)
	return touching
}

func perimeter(plot []utils.Cell) int {
	p := 0
	for _, c := range plot {
		neighbours := 0
		for _, n := range garden.Next4(c) {
			if garden.ValAt(n) == garden.ValAt(plot[0]) {
				neighbours += 1
			}
		}

		fence := 4 - neighbours
		p += fence
	}
	return p
}

func compareByRow(i, j Fence) int {
	rd := i.Row - j.Row
	if rd == 0 {
		return i.Coll - j.Coll
	}
	return rd
}

func compareByCol(i, j Fence) int {
	cd := i.Coll - j.Coll
	if cd == 0 {
		return i.Row - j.Row
	}
	return cd
}
