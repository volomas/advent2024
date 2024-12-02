package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

func isLevelSafe(prev, curr int, increasing bool) bool {
	if prev == curr || utils.Difference(prev, curr) > 3 {
		return false
	}

	if (increasing && curr < prev) || (!increasing && curr > prev) {
		return false
	}

	return true
}

func reportSafe(report []int, skipLevel bool) bool {
	reportSafe := true
	inc := report[0] < report[(len(report)-1)]
	i := 0
	j := 1
	for j < len(report) {
		prev := report[i]
		curr := report[j]
		levelSafe := isLevelSafe(prev, curr, inc)

		if levelSafe {
			i++
			j++
			continue
		}

		if !skipLevel || !reportSafe {
			// second time we hit unsafe report - we abort
			return false
		}

		if i == len(report)-1 {
			// if we comparing last pair, it doesn matter if it's unsafe, we can skip it
			return true
		}

		nextLevelSafe := isLevelSafe(prev, report[j+1], inc)
		if nextLevelSafe {
			i += 2
			j += 2
			reportSafe = false
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	safe := 0
	for _, line := range utils.ReadLines(os.Args[1]) {
		report := utils.IntFields(line)

		if reportSafe(report, false) ||
			reportSafe(report[1:], false) ||
			reportSafe(report[0:len(report)-1], false) ||
			reportSafe(report, true) {
			safe += 1
		}
	}

	fmt.Println(safe)
}
