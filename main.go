package main

import (
	"fmt";
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go from Podman!")
	})

	http.ListenAndServe(":8888", nil)
}
