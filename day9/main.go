package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"vmas/advent2024/utils"
)

var moved map[int]bool = make(map[int]bool)

func main() {
	content, _ := os.ReadFile(os.Args[1])
	input := strings.TrimSpace(string(content))

	id := 0
	memory := make([]int, 0)
	for i := 0; i < len(input); i += 2 {
		file := utils.Must(strconv.Atoi(string(input[i])))
		for j := 0; j < file; j++ {
			memory = append(memory, id)
		}

		if i == len(input)-1 {
			continue
		}
		free := utils.Must(strconv.Atoi(string(input[i+1])))
		for j := 0; j < free; j++ {
			memory = append(memory, -1)
		}

		id += 1
	}

	// fmt.Println(memory)
	// fmt.Println(Layout(memory))
	prevFile := memory[len(memory)-1]
	fileSize := 1
	for i := len(memory) - 2; i > 0; i-- {
		currentFile := memory[i]

		if prevFile == currentFile {
			fileSize += 1
			continue
		} else {
			fileToMove := prevFile
			fileToMoveSize := fileSize

			prevFile = currentFile
			fileSize = 1

			_, ok := moved[fileToMove]
			if fileToMove == -1 || ok {
				continue
			}

			// fmt.Printf("File %d with size %d. idx=%d\n", fileToMove, fileToMoveSize, i)

			s, e := FindFree(memory, fileToMoveSize, i)

			if s < 0 || e < 0 {
				// fmt.Println("No free space to move")
				continue
			}

			// fmt.Printf("Found free block [%d -> %d] of size %d\n", s, e, fileToMoveSize)

			for j := 0; j < fileToMoveSize; j++ {
				memory[s+j] = fileToMove
				memory[i+1+j] = -1
			}

			moved[fileToMove] = true
		}
	}

	checksum := 0
	for i, id := range memory {
		if id >= 0 {
			checksum += i * id
		}
	}

	// fmt.Println(Layout(memory))

	// fmt.Println("input size", len(input))
	// fmt.Println("Memory size", len(memory))
	fmt.Println(checksum)
}

func FindFree(memory []int, size int, upTo int) (int, int) {
	start, end := -1, -1

	for i := 0; i <= upTo; i++ {
		val := memory[i]

		if start == -1 && val == -1 {
			start = i
			continue
		}

		if start != -1 && val != -1 {
			end = i - 1

			if end-start+1 >= size {
				return start, end
			} else {
				start, end = -1, -1
			}
		}
	}

	if start >= 0 && end == -1 {
		if (upTo - start + 1) >= size {
			return start, upTo
		}
	}

	return -1, -1
}

func Layout(memory []int) string {
	layout := ""

	for _, e := range memory {
		if e == -1 {
			layout += string('.')
		} else {
			layout += fmt.Sprint(e)
		}
	}

	return layout
}
