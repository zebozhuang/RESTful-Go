package api

import (
    "fmt"
    "net/http"
)

type Hello struct {
}

func NewHello() *Hello {
    return new(Hello)
}

// restful api 
func (h *Hello) Get(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You launch a Get request\n")
}

func (h *Hello) Post(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You launch a Post request\n")
}

func (h *Hello) Put(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You launch a Put request\n")
}

func (h *Hello) Delete(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You lauch a Delete request\n")
}

