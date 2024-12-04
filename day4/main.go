package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

var lines []string

type Direction int

const (
	N Direction = iota
	S
	W
	E
	NW
	NE
	SW
	SE
)

var allDirections = []Direction{N, S, W, E, NW, NE, SW, SE}

func (d Direction) next(r, c int) (int, int) {
	switch d {
	case N:
		return r - 1, c
	case S:
		return r + 1, c
	case W:
		return r, c - 1
	case E:
		return r, c + 1
	case NW:
		return r - 1, c - 1
	case NE:
		return r - 1, c + 1
	case SW:
		return r + 1, c - 1
	case SE:
		return r + 1, c + 1
	default:
		return -1, -1
	}
}

func main() {
	lines = utils.ReadLines(os.Args[1])
	xmasCount := 0
	for r, line := range lines {
		for c, letter := range line {
			if letter == 'X' {
				for _, dir := range allDirections {
					if findXmas(' ', r, c, dir) {
						xmasCount += 1
					}
				}
			}
		}
	}

	fmt.Println(xmasCount)
}

// dir
func findXmas(prevLetter rune, r, c int, dir Direction) bool {
	if r < 0 || c < 0 || r >= len(lines) || c >= len(lines[0]) {
		return false
	}
	currentLetter := rune(lines[r][c])
	if currentLetter == 'S' && prevLetter == 'A' {
		return true
	}

	if r < 0 || c < 0 || r >= len(lines) || c >= len(lines[0]) {
		return false
	}

	nextR, nextC := dir.next(r, c)
	if currentLetter == 'X' && prevLetter == ' ' {
		return findXmas(currentLetter, nextR, nextC, dir)
	} else if currentLetter == 'M' && prevLetter == 'X' {
		return findXmas(currentLetter, nextR, nextC, dir)
	} else if currentLetter == 'A' && prevLetter == 'M' {
		return findXmas(currentLetter, nextR, nextC, dir)
	} else {
		return false
	}
}
