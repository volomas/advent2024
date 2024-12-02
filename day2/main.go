package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)

func main() {
	safe := 0
	for _, line := range utils.ReadLines(os.Args[1]) {
		report := utils.IntFields(line)

		reportSafe := true
		inc := report[0] < report[(len(report) - 1)]
		for i := 1; i < len(report); i++ {
			prev := report[i - 1]
			curr := report[i]

			if (utils.Difference(prev, curr) > 3)  {
				reportSafe = false
				break
			}

			if (prev == curr) {
				reportSafe = false
				break
			}

			if (inc && curr < prev) {
				reportSafe = false
				break
			} else if (!inc && curr > prev) {
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
