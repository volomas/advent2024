package utils

import (
	"log"
	"os"
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
