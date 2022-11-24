package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "Welcome to my website!")
       fmt.Fprintf(w,"Go version: %s\n", runtime.Version())
    })

    fs := http.FileServer(http.Dir("static/"))
    fs=http.StripPrefix("/static/", fs)
	http.Handle("/static/", fs)

    http.ListenAndServe(":8080", nil)
}