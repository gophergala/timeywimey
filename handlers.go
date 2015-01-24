package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	body, err := localFile("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(body))
}

func localFile(filename string) ([]byte, error) {
	page, err := ioutil.ReadFile("html/" + filename)
	if err != nil {
		return nil, err
	}
	return page, nil
}
