package web

import (
	"RESTful-Go/api"
	"RESTful-Go/handler"
)

func RegisterURL(w *Web) {

	w.Handle("/api/foo", api.NewFoo())
	w.HandleFunc("GET", "/others", api.NonRestfulApi)
	// This will serve files under http://localhost:8000/static/<filename>
	w.PathPrefix("/static/").Handler(handler.NewStaticHandler("./static", "/static/"))
}
