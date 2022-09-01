package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":3000", nil)
}

func loadFile(fileName string) (string, error) {
	bytes, error := ioutil.ReadFile(fileName)
	if error != nil {
		return "", error
	}

	return string(bytes), nil
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	var html, _ = loadFile("index.html")
	fmt.Fprintf(w, html)
}
