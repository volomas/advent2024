package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"vmas/advent2024/utils"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	tokens := 0
	for _, task := range strings.Split(string(content), "\n\n") {
		parts := strings.Split(task, "\n")
		xs := utils.IntFieldsRegex(parts[0])
		x1, x2 := xs[0], xs[1]

		ys := utils.IntFieldsRegex(parts[1])
		y1, y2 := ys[0], ys[1]

		prize := utils.IntFieldsRegex(parts[2])
		p1 := prize[0] + 10000000000000
		p2 := prize[1] + 10000000000000

		x := float64(y2*p1-y1*p2) / float64(x1*y2-x2*y1)
		y := float64(x1*p2-x2*p1) / float64(x1*y2-x2*y1)
		if x == math.Trunc(x) && y == math.Trunc(y) {
			tokens += int(x)*3 + int(y)*1
		}
	}

	fmt.Println(tokens)

}
