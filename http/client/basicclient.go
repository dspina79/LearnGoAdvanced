package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	resp, err := http.Get("http://www.google.com")

	check(err)

	defer resp.Body.Close()

	fmt.Println("Response status:", resp.StatusCode)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan() && i < 10; i++ {
		fmt.Println(scanner.Text())
	}
}
