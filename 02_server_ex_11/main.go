package main

import (
  "html/template"
  "net/http"
)

var t *template.Template

func init(){
  t = template.Must(template.ParseFiles("templates/index.html", "templates/head.html"))
}

type Page struct {
  Title, Content string
}

func displayPage(res http.ResponseWriter, req *http.Request) {
  data := &Page{
    Title : "An Example",
    Content : "Have fun stromin' da castle",
  }
  t.ExecuteTemplate(res, "index.html", data)
}

func main(){
  http.HandleFunc("/", displayPage)
  http.ListenAndServe(":5000", nil)
}
