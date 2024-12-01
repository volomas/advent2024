package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

func main() {
	left := make([]int, 0)
	right := make(map[int]int, 0)

	for _, x := range utils.ReadLines(os.Args[1]) {
		numbers := utils.IntFields(x)
		nextLeft := numbers[0]
		nextRight := numbers[1]

		left = append(left, nextLeft)
		count, ok := right[nextRight]
		if ok {
			right[nextRight] = count + 1
		} else {
			right[nextRight] = 1
		}
	}

	score := 0
	for _, num := range left {
		if count, ok := right[num]; ok {
			score += (num * count)
		}
	}
	
	fmt.Println(score)
}
