package main

import (
	"fmt"
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Printf("Server is running now!")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		spew.Dump(r)
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	var abyr ServerType
	http.ListenAndServe(":1488", nil)
}
