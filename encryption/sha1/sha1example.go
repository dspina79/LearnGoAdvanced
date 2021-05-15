package main

import (
	"crypto/sha1"
	"fmt"
)

func encrypt(s string) []byte {
	hasher := sha1.New()
	hasher.Write([]byte(s))

	octVal := hasher.Sum(nil)
	return octVal
}

func main() {
	var inputString string
	fmt.Print("Enter a value to hash: ")
	fmt.Scanf("%s", &inputString)
	encryptBytes := encrypt(inputString)
	fmt.Printf("The hashed value is: %x\n", encryptBytes)
}
