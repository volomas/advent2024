package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

type Dir rune

func (d Dir) NextPos(y, x int) (int, int) {
	switch d {
	case '^':
		return y - 1, x
	case '>':
		return y, x + 1
	case '<':
		return y, x - 1
	case 'v':
		return y + 1, x
	default:
		panic("unknown dir")
	}
}

func (d Dir) TurnRight() Dir {
	switch d {
	case '^':
		return '>'
	case '>':
		return 'v'
	case '<':
		return '^'
	case 'v':
		return '<'
	default:
		panic("unknown dir")
	}
}

var grid [][]rune

var startX, startY int
var startDir Dir
var uniqLoops map[string]bool = make(map[string]bool)

// If we tried to place '0' and it didn't work out first try, we should place it again
var tried map[string]bool = make(map[string]bool)

func IsOutAt(y, x int) bool {
	return y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1
}

func IsBlockAt(y, x int) bool {
	return grid[y][x] == '#'
}

func Pos(y, x int) string {
	return fmt.Sprintf("%d-%d", y, x)
}

func Walk(loops int, visited map[string]bool, y, x int, d Dir) int {
	visited[Pos(y, x)] = true

	nextY, nextX := d.NextPos(y, x)

	if IsOutAt(nextY, nextX) {
		return loops
	}

	newLoopCount := loops

	if IsBlockAt(nextY, nextX) {
		// if block is in front of us, turn right
		return Walk(loops, visited, y, x, d.TurnRight())
	} else if !(startX == nextX && startY == nextY) {
		key := fmt.Sprintf("%d-%d", nextY+1, nextX+1)
		// if we already placed an '0' here?
		if _, ok := tried[key]; !ok {
			// try to place block '0' in front of us, it that's not starting point
			// fmt.Printf("Placing '0' at [%d,%d] %c\n", nextY+1, nextX+1, d)
			loopVisited := make(map[string]bool)
			grid[nextY][nextX] = '#'
			loopExist := FindLoop(loopVisited, y, x, d)
			if loopExist {
				// fmt.Printf("Found loop at [%d,%d](%c)\n", nextY+1, nextX+1, d)
				uniqLoops[key] = true
			}
			tried[key] = true

			grid[nextY][nextX] = '.'
		}
	}

	return Walk(newLoopCount, visited, nextY, nextX, d)
}

func FindLoop(visited map[string]bool, y, x int, dir Dir) bool {
	key := fmt.Sprintf("%d-%d-%c", y, x, dir)

	if _, ok := visited[key]; ok {
		return true
	} else {
		visited[key] = true
	}

	nextY, nextX := dir.NextPos(y, x)

	if IsOutAt(nextY, nextX) {
		return false
	}

	if IsBlockAt(nextY, nextX) {
		return FindLoop(visited, y, x, dir.TurnRight())
	}

	return FindLoop(visited, nextY, nextX, dir)
}

func main() {
	lines := utils.ReadLines(os.Args[1])
	n := len(lines)
	m := len(lines[0])
	grid = utils.CreateGrid(n, m, '.')
	for y, line := range utils.ReadLines(os.Args[1]) {
		for x, c := range line {
			if c == '#' {
				grid[y][x] = '#'
			} else if c == '^' || c == '>' || c == '<' || c == 'v' {
				startDir = Dir(c)
				startY = y
				startX = x
			}
		}
	}
	// fmt.Printf("Starting point is [%d, %d]. start dir: %c\n", startY+1, startX+1, startDir)

	visited := make(map[string]bool)
	Walk(0, visited, startY, startX, startDir)

	fmt.Println(len(uniqLoops))
}
