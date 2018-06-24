package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	version := os.Getenv("VERSION")
	if version == "" {
		version = "DEFAULT_VERSION"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Got request")
		fmt.Fprintf(w, "Service version %s\n", version)
	})

	http.ListenAndServe(":8080", nil)
}
