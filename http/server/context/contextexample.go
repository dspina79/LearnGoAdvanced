package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkContextEnd(httpWriter http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := ctx.Err()
	fmt.Println("server", err)
	internalError := http.StatusInternalServerError
	http.Error(httpWriter, err.Error(), internalError)
}

func hello(httpWriter http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("Starging /hello call")
	defer fmt.Println("Ending /hello context")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(httpWriter, "Hello, world!\n")
	case <-time.After(20 * time.Second):
		// this wont fire because of the line above it
		fmt.Fprintf(httpWriter, "Hello again...")
	case <-ctx.Done():
		checkContextEnd(httpWriter, req)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8081", nil)
}
