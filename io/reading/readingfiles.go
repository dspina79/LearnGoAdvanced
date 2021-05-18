package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fn := "/sampletext.txt"
	dir, err := os.Getwd()
	fmt.Println(dir)
	s := dir + fn
	check(err)
	data, err := ioutil.ReadFile(s)
	check(err)

	fmt.Print(string(data))

	f, err := os.Open(s)
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	fmt.Printf("%d bytes at %d", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes at %d: %s\n", n3, o3, string(b3))

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes from the current position: %s\n", string(b4))

	f.Close()
}
