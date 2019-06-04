// Buffering a template response
// handling error
package main

import (
  "html/template"
  "net/http"
  "fmt"
  "bytes"
)

var t *template.Template

func init()  {
  t = template.Must(template.ParseFiles("templates/main.html"))
}

type Page struct {
  Title, Content string
}

func displayPage(res http.ResponseWriter, req *http.Request){
  data := &Page {
    Title : "An Example",
    Content : "Have fun stromin' da castle",
  }
  var b bytes.Buffer
  err := t.Execute(&b, data)
  if err != nil {
    fmt.Fprintf(res, "An error occured")
    return
  }
  b.WriteTo(res)
}

func main()  {
  http.HandleFunc("/", displayPage)
  http.ListenAndServe(":5000", nil)
}
