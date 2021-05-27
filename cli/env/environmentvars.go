package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("MyVar", "SimpleWord")
	fmt.Println("MYVAR:", os.Getenv("MyVar"))

	for _, e := range os.Environ() {
		envPair := strings.SplitN(e, "=", 2)
		fmt.Println(envPair[0], ": ", envPair[1])
	}
}
