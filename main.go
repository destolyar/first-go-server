package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", homeHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}

	http.ListenAndServe(":"+port, nil)
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
