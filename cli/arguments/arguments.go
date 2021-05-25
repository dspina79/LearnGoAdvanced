package main

import (
	"fmt"
	"os"
	"strconv"
)

func printArgs(args []string) {
	if len(args) > 1 {
		for _, arg := range args {
			fmt.Println(arg)
		}
	}
}

func checkArgs(args []string) bool {
	return len(args) == 3
}

func add(x, y int) int {
	return x + y
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	args := os.Args
	if checkArgs(args) {
		x, err := strconv.Atoi(args[1])
		checkError(err)
		y, err := strconv.Atoi(args[2])
		checkError(err)
		result := add(x, y)

		fmt.Printf("%d + %d = %d\n", x, y, result)

	} else {
		panic("The incorrect number of arguments were provided.")
	}
}
