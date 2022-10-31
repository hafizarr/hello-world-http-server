package main

import (
	"fmt"
	"net/http"
)

func main() {
	handlerIndex := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}

	http.HandleFunc("/", handlerIndex)

	fmt.Println("server started at localhost:4000")
	http.ListenAndServe(":9000", nil)
}
