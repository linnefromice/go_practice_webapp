package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

// http://localhost:9090/ -> "Hello my Route!"
// http://localhost:9090/xyz -> "404 page not found"
func (p *MyMux) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my Route!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}