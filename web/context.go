package web

import (
    "net/http"
    "io/ioutil"
)

type Context struct {
    Request         *http.Request
    ResponseWriter  http.ResponseWriter
    ResponseHeader  http.Header

    RawPostData     []byte
}

func NewContext(w http.ResponseWriter, r *http.Request) (*Context, error) {
    var err error

    c := new(Context)
    c.Request = r
    if w != nil {
        c.ResponseHeader = w.Header()
    }

    c.ResponseWriter = w
    c.RawPostData, err = ioutil.ReadAll(r.Body)

    r.Body.Close()

    if err != nil {
        return nil, err
    }

    return c, nil
}
