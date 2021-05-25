package main

import (
	"fmt"
	"os"
)

func checkArgs(args []string) {
	if len(args) > 1 {
		for _, arg := range args {
			fmt.Println(arg)
		}
	}
}

func main() {
	args := os.Args
	checkArgs(args)
}
