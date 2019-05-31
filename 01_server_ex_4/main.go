package main

import (
  "net/http"
  "fmt"
  "strings"
)

func main()  {
  http.HandleFunc("/hello", hello)
  http.HandleFunc("/goodbye/", goodbye)
  http.HandleFunc("/", homepage)
  http.ListenAndServe(":5000", nil)
}

func hello(res http.ResponseWriter, req *http.Request) {
  query := req.URL.Query()
  name := query.Get("name")
  if name == ""{
    name = "Manoj Chandran"
  }
  fmt.Fprint(res, "Hello, my name is ", name)
}

func goodbye(res http.ResponseWriter, req *http.Request) {
  path := req.URL.Path
  parts := strings.Split(path, "/")
  name := parts[2]
  if name == ""{
    name = "Manoj Chandran"
  }
  fmt.Fprint(res, "Goodbye ", name)
}

func homepage(res http.ResponseWriter, req *http.Request) {
  if req.URL.Path != "/"{
    http.NotFound(res, req)
  }
  fmt.Fprint(res, "The Homepage")
}
