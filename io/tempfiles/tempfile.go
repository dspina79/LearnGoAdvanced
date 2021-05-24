package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := ioutil.TempFile("", "example")
	check(err)

	fmt.Println("The temp file created is", f.Name())

	defer os.Remove(f.Name())

	_, err = f.WriteString("What a piece of work is a man\nhow noble in reason.")
	check(err)

	tempDir, err := ioutil.TempDir("", "exampleDir")
	check(err)
	defer os.RemoveAll(tempDir)

	fmt.Println("The temp directory is", tempDir)

	fname := filepath.Join(tempDir, "exampleFile2")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
	err = ioutil.WriteFile(fname, []byte("this is yet another string"), 0666)
	check(err)
}
