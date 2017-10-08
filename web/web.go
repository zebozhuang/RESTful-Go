package web

import (
	"RESTful-Go/handler"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
	"sync"
)

const (
	UNIX = "unix"
	TCP  = "tcp"
)

type Web struct {
	mux       *mux.Router
	listeners []net.Listener
}

func NewWeb() *Web {
	w := new(Web)
	w.mux = mux.NewRouter()
	return w
}

func (w *Web) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	w.mux.ServeHTTP(rw, r)
}

func (w *Web) Handle(urlpath string, hdlr interface{}) {
	h := handler.NewWebHandler(hdlr)
	w.mux.Handle(urlpath, h)
}

// Support non-restful api
func (w *Web) HandleFunc(method, urlpath string, hdlr func(http.ResponseWriter, *http.Request)) {
	h := handler.NewDummyHandler()
	h.SetFunc(method, hdlr)
	w.Handle(urlpath, h)
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
