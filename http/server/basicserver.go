package main

import (
	"fmt"
	"net/http"
)

func helloWorld(httpWriter http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(httpWriter, "hello, world!\n")
}

func fib50(httpWriter http.ResponseWriter, req *http.Request) {
	x, y := 0, 1
	for i := 0; i < 50; i++ {
		fmt.Fprintf(httpWriter, fmt.Sprint(i+1)+": "+fmt.Sprint(y)+"\n")
		tempX := x
		x = y
		y += tempX
	}
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
	http.HandleFunc("/first50fib", fib50)
	http.ListenAndServe(":8081", nil)
}
