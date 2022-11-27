package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "Hello <b>World</b>")
}

func main() {
	http.HandleFunc("/", welcome)
	http.ListenAndServe("localhost:8081", nil)
}
