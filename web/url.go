package web

import (
	"GoRestfulServer/api"
)

func RegisterURL(w *Web) {
	w.Handle("/api/foo", api.NewFoo())
	w.HandleFunc("GET", "/others", api.NonRestfulApi)
}
