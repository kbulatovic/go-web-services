package main

import (
	"net/http"
)

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yoyo"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "Hello world"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
