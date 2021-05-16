package main

import (
	"encoding/base64"
	"fmt"
)

func encode(plainText string) string {
	enc := base64.StdEncoding.EncodeToString([]byte(plainText))
	return enc
}

func decode(encoded string) string {
	decoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		panic(err)
	}

	return string(decoded)
}

func main() {
	var plainText string
	fmt.Print("Enter a value to encode: ")
	fmt.Scanf("%s", &plainText)

	encoded := encode(plainText)
	fmt.Println("The encoded value is", encoded)

	decoded := decode(encoded)
	fmt.Println("The decoded value is", decoded)
}
