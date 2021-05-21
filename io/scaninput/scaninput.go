package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	f, err := os.Create("outputString")

	if err != nil {
		panic(err)
	}

	defer f.Close()
	for scanner.Scan() {
		line := scanner.Text()
		lineLower := strings.ToUpper(line)
		if line == "exit" {
			break
		}
		f.WriteString(line + "\n")
		fmt.Println(lineLower)
	}

	f.Sync()

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
