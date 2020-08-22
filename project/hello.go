package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hello")
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/ok", okHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello!\n")
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok!\n")
}