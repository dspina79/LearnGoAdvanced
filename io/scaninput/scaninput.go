package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lineLower := strings.ToUpper(line)
		fmt.Println(lineLower)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
