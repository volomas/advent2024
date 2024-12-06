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

func IsOutAt(y, x int) bool {
	return y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[0])-1
}

func IsBlockAt(y, x int) bool {
	return grid[y][x] == '#'
}

func Pos(y, x int) string {
	return fmt.Sprintf("%d-%d", y, x)
}

func Walk(visited map[string]bool, y, x int, d Dir) map[string]bool {
	visited[Pos(y, x)] = true
	nextY, nextX := d.NextPos(y, x)

	if IsOutAt(nextY, nextX) {
		return visited
	}

	if IsBlockAt(nextY, nextX) {
		d = d.TurnRight()
		nextY, nextX = d.NextPos(y, x)
	}
	return Walk(visited, nextY, nextX, d)
}

func main() {
	var dir Dir
	lines := utils.ReadLines(os.Args[1])
	n := len(lines)
	m := len(lines[0])
	grid = utils.CreateGrid(n, m, '.')
	var posY, posX int
	for y, line := range utils.ReadLines(os.Args[1]) {
		for x, c := range line {
			if c == '#' {
				grid[y][x] = '#'
			} else if c == '^' || c == '>' || c == '<' || c == 'v' {
				dir = Dir(c)
				posY = y
				posX = x
			}
		}
	}

	visited := make(map[string]bool)
	Walk(visited, posY, posX, dir)
	fmt.Println(len(visited))
}
