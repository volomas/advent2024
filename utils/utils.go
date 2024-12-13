package utils

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var digitsRegexp *regexp.Regexp = regexp.MustCompile(`\d+`)

type Grid[T any] [][]T

func (g Grid[T]) IsOut(c Cell) bool {
	return c.Row < 0 || c.Row > len(g)-1 || c.Coll < 0 || c.Coll > len(g[0])-1
}

func (g Grid[T]) ValAt(c Cell) T {
	return g[c.Row][c.Coll]
}

func (g Grid[T]) Next4(c Cell) []Cell {
	next := make([]Cell, 0)

	up := Cell{c.Row - 1, c.Coll}
	down := Cell{c.Row + 1, c.Coll}
	left := Cell{c.Row, c.Coll - 1}
	right := Cell{c.Row, c.Coll + 1}

	for _, n := range []Cell{up, down, left, right} {
		if !g.IsOut(n) {
			next = append(next, n)
		}
	}

	return next
}

func (g Grid[T]) Next8(c Cell) []Cell {
	next := g.Next4(c)
	for _, n := range []Cell{
		{c.Row - 1, c.Coll - 1},
		{c.Row - 1, c.Coll + 1},
		{c.Row + 1, c.Coll - 1},
		{c.Row + 1, c.Coll + 1},
	} {
		if !g.IsOut(n) {
			next = append(next, n)
		}
	}

	return next
}

type Cell struct{ Row, Coll int }

func NewCell(row, coll int) Cell {
	return Cell{row, coll}
}

func ParseIntGrid(lines []string) Grid[int] {
	grid := make([][]int, len(lines))
	for row, line := range lines {
		grid[row] = make([]int, len(line))
		for col, val := range line {
			intVal := Must(strconv.Atoi(string(val)))
			grid[row][col] = intVal
		}
	}
	return grid
}

func ParseCharGrid(lines []string) Grid[rune] {
	grid := make([][]rune, len(lines))
	for row, line := range lines {
		grid[row] = make([]rune, len(line))
		for col, val := range line {
			grid[row][col] = val
		}
	}
	return grid
}

func ReadLines(fileName string) []string {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Failed to read file", fileName, err)
	}

	content := string(fileBytes)

	return strings.Split(strings.TrimSpace(content), "\n")
}

// Creates NxM grid with all cells set to initialValue
func CreateGrid[T any](n, m int, initialValue T) [][]T {
	grid := make([][]T, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]T, m)
		for j := 0; j < m; j++ {
			grid[i][j] = initialValue
		}
	}
	return grid
}

// Prints 2d grid
func PrintGrid[T any](grid [][]T) {
	for _, x := range grid {
		for _, y := range x {
			fmt.Printf("%c", y)
		}
		fmt.Println()
	}
}

func Int64Fields(str string) []int64 {
	fields := strings.Fields(str)
	res := make([]int64, len(fields))
	for i, f := range fields {
		num := Must(strconv.ParseInt(f, 10, 64))
		res[i] = num
	}

	return res
}

func IntFields(str string) []int {
	fields := strings.Fields(str)
	res := make([]int, len(fields))
	for i, f := range fields {
		num := Must(strconv.Atoi(f))
		res[i] = num
	}

	return res
}

func IntFieldsRegex(str string) []int {
	digits := digitsRegexp.FindAllString(str, -1)
	res := make([]int, len(digits))
	for i, f := range digits {
		num := Must(strconv.Atoi(f))
		res[i] = num
	}

	return res
}

func Must[T any](obj T, err error) T {
	if err != nil {
		panic(err)
	}
	return obj
}

func Difference(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
