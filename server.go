package main

import (
    "RESTful-Go/web"
)

func main() {

    var err error

    w := web.NewWeb()     

    web.RegisterURL(w)

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
