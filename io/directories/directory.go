package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	err := os.Mkdir("subdir", 0755)
	check(err)

	defer os.RemoveAll("subdir")

	createEmptyFile := func(fileName string) {
		d := []byte("")
		check(ioutil.WriteFile(fileName, d, 0644))
	}

	createEmptyFile("subdir/file1.txt")
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2.txt")
	createEmptyFile("subdir/parent/file3.txt")
	createEmptyFile("subdir/parent/child/file4.txt")

	c, err := ioutil.ReadDir("subdir/parent")
	check(err)

	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}
}
