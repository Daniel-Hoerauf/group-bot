package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
	fmt.Println("Content received!")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}