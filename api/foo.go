package api

import (
	"fmt"
	"net/http"
)

type Foo struct {
}

func NewFoo() *Foo {
	return new(Foo)
}

// restful api
func (h *Foo) Get(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You launch a Get request\n")
}

func (h *Foo) Post(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You launch a Post request\n")
}

func (h *Foo) Put(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You launch a Put request\n")
}

func (h *Foo) Delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You launch a Delete request\n")
}

// func NonRestfulApi(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "You launch a Non restful api.\n")
// }
