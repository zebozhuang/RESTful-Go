package main

import (
    "fmt"
    "net/http"
    "GoRestfulServer/web"
)

type Hello struct {
}

func NewHello() *Hello {
    return new(Hello)
}

func (h *Hello) Get(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {

    var err error

    w := web.NewWeb()     
    w.Handle("/hello", NewHello())

    err =  w.Listen("tcp", ":8000")
    if err != nil {
        println(err)
    }

    err = w.Listen("unix", "/tmp/app.sock")
    if err != nil {
        println(err)
    }
    w.Serve()
}
