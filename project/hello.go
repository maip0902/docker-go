package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!\n")
}