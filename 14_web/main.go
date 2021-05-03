package main

import (
	"fmt"
	"net/http"

)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Mars!</h1>")
}

func main () {
	http.HandleFunc("/", index)
	http.ListenAndServe(":3000", nil)
}