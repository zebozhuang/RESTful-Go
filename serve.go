package main

import (
    "fmt"
    "net/http"
    "GoRestfulServer/web"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {

    var err error

    w := web.NewWeb()     
    w.Handle("/abc", handler)

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
