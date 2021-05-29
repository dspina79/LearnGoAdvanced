package main

import (
	"fmt"
	"net/http"
)

func helloWorld(httpWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(httpWriter, "hello, world!\n")
}
