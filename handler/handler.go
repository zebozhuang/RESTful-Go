package handler 

import (
    "fmt"
    "net/http"
)
const (
    POST = "POST"
    GET = "GET"
    DELETE = "DELETE"
    PUT = "PUT"
    HEAD = "HEAD"
)

type Getter interface {
    Get(http.ResponseWriter, *http.Request)
}

type Poster interface {
    Post(http.ResponseWriter, *http.Request)
}

type Putter interface {
    Put(http.ResponseWriter, *http.Request)
}

type Deleter interface {
    Delete(http.ResponseWriter, *http.Request)
}

type Header interface {
    Head(http.ResponseWriter, *http.Request)
}

type WebHandler struct {
    getter Getter
    poster Poster
    putter Putter
    deleter Deleter
    header Header
}

func NewWebHandler(handler interface{}) *WebHandler {
    h := new(WebHandler)
    h.getter, _ = handler.(Getter)
    h.poster, _  = handler.(Poster)
    h.putter, _  = handler.(Putter)
    h.deleter, _  = handler.(Deleter)
    h.header, _ = handler.(Header)

    if h.getter == nil && h.poster == nil && 
        h.putter == nil && h.deleter == nil && 
        h.header == nil {
        return nil
    }
    return h
}

func (h *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    if r.Method == GET && h.getter != nil {
        h.getter.Get(w, r)
    } else if r.Method == POST && h.poster != nil {
        h.poster.Post(w, r)
    } else if r.Method == DELETE && h.deleter != nil {
        h.deleter.Delete(w, r)
    } else if r.Method == PUT && h.putter!= nil {
        h.putter.Put(w, r)
    } else if r.Method == HEAD && h.header != nil {
        h.header.Head(w, r)
    } else {
        h.methodNotAllowed(w, r) 
    }
}

func (h *WebHandler) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "method not allowed")
}
