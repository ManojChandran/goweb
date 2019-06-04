package main

import (
  "html/template"
  "net/http"
)

type Page struct{
  Title, Content string
}

func displayPage(res http.ResponseWriter, req *http.Request) {
  data := &Page {
    Title : "An Example",
    Content : "Have fun stromin' da castle",
  }
  t := template.Must(template.ParseFiles("templates/main.html"))
  t.Execute(res, data)
}

func main()  {
  http.HandleFunc("/",displayPage)
  http.ListenAndServe(":5000", nil)
}
