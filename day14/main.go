package main

import (
	"fmt"
	"os"
	"slices"
	"vmas/advent2024/utils"
)

const (
	//sample
	// H = 7
	// W = 11

	//input
	H = 103
	W = 101
)

type Robot struct {
	x  int
	y  int
	dx int
	dy int
}

func (r Robot) String() string {
	return fmt.Sprintf("[%d,%d]", r.x, r.y)
}

func (r *Robot) Move() {
	r.x = (r.x + r.dx + W) % W
	r.y = (r.y + r.dy + H) % H
}

func main() {
	lines := utils.ReadLines(os.Args[1])
	robots := make([]*Robot, len(lines))
	for i, line := range lines {
		numbers := utils.IntFieldsRegex(line)
		robots[i] = &Robot{numbers[0], numbers[1], numbers[2], numbers[3]}
	}

	for i := 1; i <= 10000; i++ {
		for _, r := range robots {
			r.Move()
		}

		if HasTree(robots) {
			Plot(robots)
			fmt.Println(i)
			break
		}
	}
}

func Plot(robots []*Robot) {
	grid := utils.CreateGrid(H, W, '.')

	for _, r := range robots {
		grid[r.y][r.x] = '*'
	}

	utils.PrintGrid(grid)
}

// Checks if has straight line with 10 robots
func HasTree(robots []*Robot) bool {
	slices.SortFunc(robots, func(r1, r2 *Robot) int {
		rd := r1.y - r2.y
		if rd == 0 {
			return r1.x - r2.x
		}
		return rd
	})

	lineCount := 0
	for i := 1; i < len(robots); i++ {
		prev := robots[i-1]
		curr := robots[i]

		if prev.y == curr.y && prev.x == curr.x-1 {
			lineCount += 1
		} else if prev.y == curr.y && prev.x == curr.x {
			continue
		} else {
			lineCount = 0
		}

		if lineCount == 10 {
			return true
		}
	}

	return false
}
