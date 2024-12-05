package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(fileName string) []string {
	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal("Failed to read file", fileName, err)
	}

	content := string(fileBytes)

	return strings.Split(strings.TrimSpace(content), "\n")
}

//Creates NxM grid with all cells set to initialValue	
func CreateGrid[T any](n, m int, initialValue T) [][]T {
	grid := make([][]T, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]T, m)
		for j := 0; j < m; j++ {
			grid[i][j] = initialValue
		}
	}
	return grid
}

//Prints 2d grid
func PrintGrid[T any](grid [][]T) {
	for _, x := range grid {
		for _, y := range x {
			fmt.Printf("%v ", y)
		}
		fmt.Println()
	}
}

func IntFields(str string) []int {
	fields := strings.Fields(str)
	res := make([]int, len(fields))
	for i, f := range fields {
		num := Must(strconv.Atoi(f))
		res[i] = num
	}

	return res
}

func IntFieldsSep(str, sep string) []int {
	fields := strings.Split(str, sep)
	res := make([]int, len(fields))
	for i, f := range fields {
		num := Must(strconv.Atoi(f))
		res[i] = num
	}

	return res
}

func Must[T any](obj T, err error) T {
    if err != nil {
        panic(err)
    }
    return obj
}

func Difference(a, b int) int {
	if a < b {
		return b - a
	} else {
		return a - b
	}
}
