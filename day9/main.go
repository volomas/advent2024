package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"vmas/advent2024/utils"
)

func main() {
	file, _ := os.ReadFile(os.Args[1])
	input := strings.TrimSpace(string(file))

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

	fmt.Println(memory)
	// fmt.Println(Layout(memory))
	s, e := FirstFreeBlock(memory)
	fmt.Printf("%d-%d\n", s, e)
	for i := len(memory) - 1; i > 0; i-- {
		s, _ := FirstFreeBlock(memory)
		if s == -1 || s > i {
			break
		}
		if memory[i] != -1 {
			memory[s] = memory[i]
			memory[i] = -1
		}
	}

	checksum := 0
	for i, id := range memory {
		if id >= 0 {
			checksum += i * id
		}
	}

	fmt.Println(Layout(memory))
	fmt.Println(checksum)
}

func FirstFreeBlock(memory []int) (int, int) {
	start, end := -1, -1
	for i, val := range memory {
		if start == -1 && val == -1 {
			start = i
		}

		if start != -1 && val != -1 {
			end = i - 1
			break
		}
	}
	return start, end
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
