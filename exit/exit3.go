package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("I see you shivver with anti..ci...")
	// the following line will never run
	defer fmt.Println("...pation!")
	<-time.After(3 * time.Second)
	os.Exit(3)
}
