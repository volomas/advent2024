package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"vmas/advent2024/utils"
)

func main() {
	bytes, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	content := string(bytes)
	do := true
	sum := 0

	for i, char := range content {
		if char == 'd' {
			enable, ok := parseEnableAt(content, i)
			if (ok) {
				do = enable
			}
		}

		if do && char == 'm' {
			x, y := parseMulAt(content, i)
			sum += x * y
		}
	}
	fmt.Println(sum)
}

// parse don't() or do()
func parseEnableAt(content string, idx int) (bool, bool) {
	maybeDontOrDo := string(content[idx:idx+7])
	if (strings.Contains(maybeDontOrDo, "don't()")) {
		return false, true
	} else if (strings.Contains(maybeDontOrDo, "do()")) {
		return true, true
	}
	return false, false
}

// parse mul(x,y)
func parseMulAt(content string, idx int) (int, int) {
	mul := utils.Must(regexp.Compile(`^mul\([0-9]+,[0-9]+\)`))
	found := mul.FindString(content[idx:])
	if found == "" {
		return 0, 0
	} else {
		found, _ = strings.CutPrefix(found, "mul(")
		found, _ = strings.CutSuffix(found, ")")
		args := strings.Split(found, ",")
		return utils.Must(strconv.Atoi(args[0])), utils.Must(strconv.Atoi(args[1]))
	}
}
