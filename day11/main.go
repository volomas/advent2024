package main

import (
	"fmt"
	"os"
	"strconv"
	"vmas/advent2024/utils"
)

const BLINKS = 75

func main() {
	lines := utils.ReadLines(os.Args[1])
	stones := utils.IntFields(lines[0])

	//count occurence of each stone (refresh at each step)
	_map := make(map[int]int)

	for _, s := range stones {
		_map[s]++
	}

	for i := 1; i <= BLINKS; i++ {
		newMap := make(map[int]int)
		for stone, count := range _map {
			for _, newStone := range split(stone) {
				newMap[newStone] += count
			}
		}
		_map = newMap
	}

	count := 0
	for _, v := range _map {
		count += v
	}
	fmt.Println(count)
}

func split(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	strNum := fmt.Sprintf("%d", stone)
	if len(strNum)%2 == 0 {
		left := utils.Must(strconv.Atoi(strNum[0 : len(strNum)/2]))
		right := utils.Must(strconv.Atoi(strNum[len(strNum)/2:]))
		return []int{left, right}
	}

	return []int{stone * 2024}
}
