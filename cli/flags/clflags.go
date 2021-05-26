package main

import (
	"flag"
	"fmt"
)

func main() {

	fnPtr := flag.String("fn", "john", "First Name")
	lnPtr := flag.String("ln", "smith", "lastName Name")
	readOnly := flag.Bool("readOnly", false, "A boolean")
	var svar string
	flag.StringVar(&svar, "svar", "foo", "A string variable")

	flag.Parse()

	fmt.Println("First Name", *fnPtr)
	fmt.Println("Last Name", *lnPtr)
	fmt.Println("Read Only", *readOnly)
	fmt.Println("svar: ", svar)
}
