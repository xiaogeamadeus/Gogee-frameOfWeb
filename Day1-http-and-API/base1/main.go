package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// set 2 route- "/" and "/hello"
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// launch web service
	// :9999 : address AT 9999 port
	// nil : solve instance by Standard library
	log.Fatal(http.ListenAndServe(":9999", nil))
}

//handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

//handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
