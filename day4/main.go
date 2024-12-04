package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

var lines []string

func main() {
	lines = utils.ReadLines(os.Args[1])
	xmasCount := 0
	for r := 2; r < len(lines); r++ {
		for c := 2; c < len(lines[0]); c++ {
			if isXmas(r, c) {
				xmasCount += 1
			}
		}
	}

	fmt.Println(xmasCount)
}

func isXmas(r, c int) bool {
	center := rune(lines[r-1][c-1])
	if center != 'A' {
		return false
	}

	a1 := rune(lines[r-2][c-2])
	a2 := rune(lines[r][c])

	b1 := rune(lines[r-2][c])
	b2 := rune(lines[r][c-2])

	d1 := (a1 == 'M' && a2 == 'S') || (a1 == 'S' && a2 == 'M')
	d2 := (b1 == 'M' && b2 == 'S') || (b1 == 'S' && b2 == 'M')

	return d1 && d2
}
