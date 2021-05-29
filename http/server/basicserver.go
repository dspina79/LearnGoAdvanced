package main

import (
	"fmt"
	"net/http"
)

func helloWorld(httpWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(httpWriter, "hello, world!\n")
}

func printHeaders(httpWriter http.ResponseWriter, req *http.Request) {
	for name, header := range req.Header {
		for _, h := range header {
			fmt.Fprintf(httpWriter, "%v: %s", name, h)
		}
	}
}

func main() {
	http.HandleFunc("/hello", helloWorld)
	http.HandleFunc("/headers", printHeaders)

	http.ListenAndServe(":8081", nil)
}
