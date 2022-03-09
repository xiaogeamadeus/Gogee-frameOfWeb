package Gogee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by Gogee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct{
	router map[string]HandlerFunc
}

// New is the constructor of Gogee.Engine
func New() *Engine{
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc){
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc){
	engine.addRoute("GET", pattern, handler)
}

//POST defines the method to POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc){
	engine.addRoute("POST", pattern, handler)
}

//Run definesd the method to start a http server
func (engine *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr, engine)
}

// To use the ListenAndServe, we need to set a ServeHTTP struct
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	}else {
		fmt.Fprintf(w, "404 MOT FOUND: %s\n", req.URL)
	}
}