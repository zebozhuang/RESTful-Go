package handler

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	POST   = "POST"
	GET    = "GET"
	DELETE = "DELETE"
	PUT    = "PUT"
	HEAD   = "HEAD"
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
	getter  Getter
	poster  Poster
	putter  Putter
	deleter Deleter
	header  Header
}

func NewWebHandler(handler interface{}) *WebHandler {
	h := new(WebHandler)
	h.getter, _ = handler.(Getter)
	h.poster, _ = handler.(Poster)
	h.putter, _ = handler.(Putter)
	h.deleter, _ = handler.(Deleter)
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
	} else if r.Method == PUT && h.putter != nil {
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

type DummyHandler struct {
	getter  func(w http.ResponseWriter, r *http.Request)
	poster  func(w http.ResponseWriter, r *http.Request)
	putter  func(w http.ResponseWriter, r *http.Request)
	deleter func(w http.ResponseWriter, r *http.Request)
	header  func(w http.ResponseWriter, r *http.Request)
}

func NewDummyHandler() *DummyHandler {
	h := new(DummyHandler)
	h.getter = h.methodNotAllowed
	h.poster = h.methodNotAllowed
	h.putter = h.methodNotAllowed
	h.deleter = h.methodNotAllowed
	return h
}

func (h *DummyHandler) SetFunc(method string, handler func(http.ResponseWriter, *http.Request)) {
	switch strings.ToUpper(method) {
	case GET:
		h.getter = handler
	case POST:
		h.poster = handler
	case PUT:
		h.putter = handler
	case DELETE:
		h.deleter = handler
	case HEAD:
		h.header = handler
	default:
		panic("Method not allowed")
	}
}

func (h *DummyHandler) Get(w http.ResponseWriter, r *http.Request) {
	h.getter(w, r)
}

func (h *DummyHandler) Post(w http.ResponseWriter, r *http.Request) {
	h.poster(w, r)
}

func (h *DummyHandler) Put(w http.ResponseWriter, r *http.Request) {
	h.putter(w, r)
}

func (h *DummyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	h.deleter(w, r)
}

func (h *DummyHandler) Header(w http.ResponseWriter, r *http.Request) {
	h.header(w, r)
}

func (h *DummyHandler) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "method not allowed")
}
