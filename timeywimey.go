package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("", nil)
	http.HandleFunc("/", DefaultHandler)

	log.Fatal(http.ListenAndServe(address(), nil))
}

// Retrieve the web server address from the environment variable TW_SERVER.
// If the environment variable is not set then default to `localhost:8080`.
func address() string {
	env := os.Getenv("TW_SERVER")
	if env == "" {
		return "localhost:8080"
	}
	return env
}
