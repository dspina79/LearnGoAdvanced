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
	fn := "/newfile.txt"
	wd, err := os.Getwd()
	check(err)
	fullPath := wd + fn

	content := []byte("This is the file\nthat includes the next level\nof data.")
	err = ioutil.WriteFile(fullPath, content, 0644)
	check(err)

	f, err := os.Create(wd + "/newfile2.txt")
	check(err)

	defer f.Close()

	content2 := []byte("What a piece of work is a man\nhow noble in reason\n")
	data2, err := f.Write(content2)
	fmt.Printf("The file was written with %d bytes", data2)

	data3, err := f.WriteString("this should appear as a third line\n")
	fmt.Printf("Write another %d bytes", data3)

}
