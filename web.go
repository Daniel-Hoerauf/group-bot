package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
	r.ParseForm()
	fmt.Println("Content received!")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	message := string(body[:])
	fmt.Println(message)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":80", nil)
}
