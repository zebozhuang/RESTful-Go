package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "net"
    "net/http"
    "os"
    "sync"
)

const (
    UNIX = "unix"
    TCP = "tcp"
)

type Web struct {
    mux *mux.Router 
    listeners []net.Listener
}

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

func NewWeb() *Web{
    w := new(Web)
    w.mux = mux.NewRouter()
    return w
}

func (w *Web) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
    w.mux.ServeHTTP(rw, req)  
}

func (w *Web) Handle(urlpath string, handler func(http.ResponseWriter, *http.Request)) {
    h := NewWebHandler(handler)

    w.mux.Handle(urlpath, h)
}

func (w *Web) Listen(protocol, addr string) error {
    if protocol == UNIX {
        os.Remove(addr)
    }

    l, err := net.Listen(protocol, addr) 
    if err != nil {
        return err
    }

    if protocol == UNIX {
        os.Chmod(addr, 0666)
    }

    w.listeners = append(w.listeners, l) 
    return nil
}

func (w *Web) Serve() error {
    var err error

    if len(w.listeners) == 0 {
        return fmt.Errorf("no listener")    
    }

    wg := sync.WaitGroup{}
    server := func(l net.Listener) {
        svr := http.Server{Handler: w}

        err = svr.Serve(l)
        l.Close()

        wg.Done()     
    }

    for _, l := range w.listeners {
        wg.Add(1)
        go server(l)
    }
    wg.Wait()
    return err
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}


func main() {

    var err error

    w := NewWeb()     
    w.Handle("/abc", handler)

    err =  w.Listen(TCP, ":8000")
    if err != nil {
        println(err)
    }

    err = w.Listen(UNIX, "/tmp/app.sock")
    if err != nil {
        println(err)
    }
    w.Serve()
}
