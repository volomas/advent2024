package main

import (
	"fmt"
	"os"
	"strconv"
	"vmas/advent2024/utils"
)

const BLINKS = 75

var table []map[int64]int = make([]map[int64]int, BLINKS)

func main() {
	for i := 0; i < BLINKS; i++ {
		table[i] = make(map[int64]int)
	}

	lines := utils.ReadLines(os.Args[1])
	stones := utils.Int64Fields(lines[0])

	count := 0
	for _, stone := range stones {
		count += Count(BLINKS, stone)
	}
	fmt.Println(count)
}

func Count(blinks int, stone int64) int {
	stones := split(stone)
	if blinks == 1 {
		return len(stones)
	} else if c, ok := table[blinks-1][stone]; ok {
		return c
	} else {
		count := 0
		for _, stone := range stones {
			count += Count(blinks-1, stone)
		}

		table[blinks-1][stone] = count
		return count
	}

}

func split(stone int64) []int64 {
	if stone == 0 {
		return []int64{1}
	} else {
		strNum := fmt.Sprintf("%d", stone)
		if len(strNum)%2 == 0 {
			left := utils.Must(strconv.ParseInt(strNum[0:len(strNum)/2], 10, 64))
			right := utils.Must(strconv.ParseInt(strNum[len(strNum)/2:], 10, 64))
			return []int64{left, right}
		} else {
			return []int64{stone * 2024}
		}
	}
}
