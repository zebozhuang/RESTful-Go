package api

import (
	"html/template"
	//    "os"
	// "fmt"
	"net/http"
)

type Actor struct {
	UserName string
}

//func NonRestfulApi(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "You launch a Non restful api.\n")
//}

func NonRestfulApi(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/index.html")
	t.Execute(w, "Hello World")
	// actor := Actor{UserName: "Bob@gmail.com"}

	//    t.Execute(os.Stdout, actor)
	// t.Execute(w, actor)
}
