package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(httpWriter http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	fmt.Println("Starging /hello call")
	defer fmt.Println("Ending /hello context")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(httpWriter, "Hello, world!")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server", err)
		internalError := http.StatusInternalServerError
		http.Error(httpWriter, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8081", nil)
}
