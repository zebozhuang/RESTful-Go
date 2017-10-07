package api

import (
    "html/template"
//    "os"
    "fmt"
    "net/http"
)

type Actor struct {
    UserName string
}


//func NonRestfulApi(w http.ResponseWriter, r *http.Request) {
//    fmt.Fprintf(w, "You launch a Non restful api.\n")
//}

func TestTemplate(w http.ResponseWriter, r *http.Request) {
    t := template.New("map data demo template")
    t, _ = t.Parse("Hello, {{.UserName}}!\n")

    actor := Actor{UserName: "Bob@gmail.com"}

//    t.Execute(os.Stdout, actor)
    t.Execute(w, actor)
}

func TestTemplate2(w http.ResponseWriter, r *http.Request) {
    actor := Actor{UserName: "Bob"}
    t, err := template.ParseFiles("template/tmpl.html")
    if err != nil {
        fmt.Println("Error:", err.Error())
    }
    err = t.Execute(w, actor)
    if err != nil {
        fmt.Println("Error:", err.Error())
    }
}
