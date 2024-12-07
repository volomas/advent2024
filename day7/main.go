package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"vmas/advent2024/utils"
)

func main() {
	var sum int64 = 0
	for _, line := range utils.ReadLines(os.Args[1]) {
		parts := strings.Split(line, ": ")
		target, _ := strconv.ParseInt(parts[0], 10, 64)
		nums := utils.IntFields(parts[1])
		if Eval(target, nums, 1, int64(nums[0])) {
			sum += target
		}
	}
	fmt.Println(sum)

}

func Eval(target int64, nums []int, pos int, acc int64) bool {
	finished := pos == len(nums)
	if finished {
		return target == acc
	}

	nextNum := nums[pos]

	return Eval(target, nums, pos+1, acc+int64(nextNum)) ||
		Eval(target, nums, pos+1, acc*int64(nextNum)) ||
		Eval(target, nums, pos+1, Concat(acc, int64(nextNum)))
}

func Concat(a, b int64) int64 {
	astr := strconv.FormatInt(a, 10)
	bstr := strconv.FormatInt(b, 10)

	res, err := strconv.ParseInt(astr+bstr, 10, 64)
	if err != nil {
		panic(err)
	}
	return res
}
