package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is the uni handler for all requests
type Engine struct{}

// ResponseWriter: create the response of this HTTP request.
// http.Request: all the information of this HTTP request. For example, the address, Header and Body

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	switch req.URL.Path{
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header{
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	
	// Transfer all the request to our own solution.
	engine := new(Engine)
	// launch web service
	// :9999 : address AT 9999 port
	// 
	log.Fatal(http.ListenAndServe(":9999", engine))
}
