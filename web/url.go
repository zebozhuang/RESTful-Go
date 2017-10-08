package web

import (
    "net/http"
	"RESTful-Go/api"
)

func RegisterURL(w *Web) {
    w.Handle("/css/", http.FileServer(http.Dir("template"))) 
    w.Handle("/js/", http.FileServer(http.Dir("template"))) 
    w.Handle("/", http.FileServer(http.Dir("template"))) 
	w.Handle("/api/foo", api.NewFoo())
	w.HandleFunc("GET", "/others", api.NonRestfulApi)
    w.HandleFunc("GET", "/template", api.TestTemplate)
    w.HandleFunc("GET", "/template2", api.TestTemplate2)
}
