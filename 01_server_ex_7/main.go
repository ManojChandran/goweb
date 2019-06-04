// Extending templates with funcions
package main

import (
  "net/http"
  "html/template"
  "time"
)

var tpl = `<!DOCTYPE HTML>
              <html>
                <head>
                  <meta charset="utf-8">
                  <title>{{.Title}}</title>
                </head>
                <body>
                  <h1>{{.Title}}</h1>
                  <p>{{ .Date | dateFormat "Jan 2, 2006" }}</p>
                </body>
              </html>`

var funcMap = template.FuncMap {
  "dateFormat": dateFormat,
}

func dateFormat(layout string, d time.Time) string  {
  return d.Format(layout)
}

func serveTemplate (res http.ResponseWriter,req *http.Request) {
  t := template.New("date")
  t.Funcs(funcMap)
  t.Parse(tpl)
  data := struct { Title string
                   Date time.Time }{
    Title: "Hello",
    Date: time.Now(),
  }
  t.Execute(res, data)
}

func main() {
  http.HandleFunc("/", serveTemplate)
  http.ListenAndServe(":5000", nil)
}
