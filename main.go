package main

import (
	"fmt"
	"os"
	"os/exec"
)

// simple runner: go run . <DAY_NUMBER> [input file name]
func main() {
	day := os.Args[1]
	var cmd *exec.Cmd
	if len(os.Args) > 2 {
		cmd = exec.Command("go", "run", "./day"+day, "day"+day+"/" + os.Args[2] + ".txt")
	} else {
		cmd = exec.Command("go", "run", "./day"+day, "day"+day+"/input.txt")
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error ", err)
	}
}
