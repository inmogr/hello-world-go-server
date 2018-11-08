package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}

func foo(responseWritter http.ResponseWriter, request *http.Request) {
	io.WriteString(responseWritter, "Hello World")
}
