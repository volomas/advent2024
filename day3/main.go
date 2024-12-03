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

	mul, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		panic(err)
	}
	muls := mul.FindAllString(content, -1)
	sum := 0
	for _, m := range muls {
		args := strings.Split(m, ",")
		a1, _ := strings.CutPrefix(args[0], "mul(")
		a2, _ := strings.CutSuffix(args[1], ")")
		n1 := utils.Must(strconv.Atoi(a1))
		n2 := utils.Must(strconv.Atoi(a2))
		sum += n1 * n2
	}
	fmt.Println(sum)
}
