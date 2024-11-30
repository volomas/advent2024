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

	grid := utils.CreateGrid(10, 10, "*")
	utils.PrintGrid(grid)
	grid2 := utils.CreateGrid(10, 20, "x")
	utils.PrintGrid(grid2)
}
