package web

import (
    "GoRestfulServer/api"
)

func RegisterURL(w *Web) {
    w.Handle("/api/hello", api.NewHello())
}
