package main

import (
	"fmt"
	"os"
	"vmas/advent2024/utils"
)


func main() {
	for idx, x := range utils.ReadLines(os.Args[1]) {
		fmt.Println("i", idx, "line", x)
	}
}
