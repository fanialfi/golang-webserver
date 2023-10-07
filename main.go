package main

import (
	"fmt"
	"net/http"

	"github.com/fanialfi/golang-webserver/routing"
)

var port string = ":8080"

func main() {
	http.HandleFunc("/index", routing.Index)
	http.HandleFunc("/home", routing.Home)
	http.HandleFunc("/about", routing.About)
	http.HandleFunc("/contact", routing.Contact)
	http.HandleFunc("/", routing.Root)

	http.HandleFunc("/bla", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	fmt.Printf("starting web server at localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
