package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

func isUnsafe(prev, curr int, increasing bool) bool {
	if prev == curr || utils.Difference(prev, curr) > 3 {
		return true
	}

	if (increasing && curr < prev) || (!increasing && curr > prev) {
		return true
	}

	return false
}

func main() {
	safe := 0
	for _, line := range utils.ReadLines(os.Args[1]) {
		report := utils.IntFields(line)

		reportSafe := true
		inc := report[0] < report[(len(report)-1)]
		for i := 1; i < len(report); i++ {
			prev := report[i-1]
			curr := report[i]

			if isUnsafe(prev, curr, inc) {
				reportSafe = false
				break
			}

		}

		if reportSafe {
			safe += 1
		}
	}

	fmt.Println(safe)
}
