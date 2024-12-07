package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"vmas/advent2024/utils"
)

var order map[int][]int = make(map[int][]int)

func compare(a, b int) int {
	afterA := order[a]
	for _, x := range afterA {
		if x == b {
			return -1
		}
	}
	afterB := order[b]
	for _, x := range afterB {
		if x == a {
			return 1
		}
	}
	return 0
}

func buildOrder(pairsBlock string) {
	for _, pair := range strings.Split(pairsBlock, "\n") {
		numbers := strings.Split(pair, "|")
		n1 := utils.Must(strconv.Atoi(numbers[0]))
		n2 := utils.Must(strconv.Atoi(numbers[1]))
		arr, ok := order[n1]
		if ok {
			order[n1] = append(arr, n2)
		} else {
			order[n1] = []int{n2}
		}
	}
}

func middle(list []int) int {
	// always odd
	return list[len(list)/2]
}

func main() {
	f := utils.Must(os.ReadFile(os.Args[1]))
	content := string(f)
	parts := strings.Split(content, "\n\n")
	buildOrder(parts[0])

	sum := 0
	for _, csv := range strings.Split(strings.TrimSpace(parts[1]), "\n") {
		list := utils.IntFieldsSep(csv, ",")
		if !slices.IsSortedFunc(list, compare) {
			slices.SortFunc(list, compare)
			sum += middle(list)
		}
	}

	fmt.Println(sum)
}
