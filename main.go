package main

import (
	"fmt"
	"net/http"
)

var port string = ":8080"

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hallo")
	})

	http.HandleFunc("/index", index)

	fmt.Printf("starting web server at localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
