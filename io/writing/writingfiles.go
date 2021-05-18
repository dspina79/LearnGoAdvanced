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

func writeDataToFile(data string, filePath string) (int, error) {

	f, err := os.Create(filePath)
	check(err)

	defer f.Close()
	resultInt, err := f.WriteString(data)
	f.Sync()
	return resultInt, err
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
	fmt.Printf("The file was written with %d bytes\n", data2)

	data3, err := f.WriteString("this should appear as a third line\n")
	fmt.Printf("Wrote another %d bytes\n", data3)

	// submit the information
	f.Sync()
	writeDataToFile("this is example text", wd+"/myfile.txt")
}
