package handler 

import (
    "net/http"
)

type WebHandler struct {
    handler func(http.ResponseWriter, *http.Request)
}

func NewWebHandler(handler func(http.ResponseWriter, *http.Request)) *WebHandler {
    h := new(WebHandler)
    h.handler = handler
    return h
}

func (h *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    h.handler(w, r)
}
