package main

import (
  "fmt"
  "net/http"
  "path"
  "strings"
)

func main()  {
  pr := newPathResolver()
  pr.Add("GET /hello",hello)
  pr.Add("* /goodbye/*", goodbye)
//  fmt.Println(pr)
}

func newPathResolver() *pathResolver {
  return &pathResolver{make(map[string]http.HandlerFunc)}
}

type pathResolver struct {
  handlers map[string]http.HandlerFunc
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc)  {
  p.handlers[path] = handler
}

func (p *pathResolver) ServerHTTP(res http.ResponseWriter, req *http.Request)  {
  check := req.Method + " " + req.URL.Path
  for pattern, handlerFunc := range p.handlers{
    if ok, err := path.Match(pattern, check); ok && err == nil{
      handlerFunc(res, req)
      return
    } else if err != nil {
      fmt.Fprint(res,req)
    }
  }
  http.NotFound(res, req)
}

func hello(res http.ResponseWriter, req *http.Request)  {
  query := req.URL.Query()
  name := query.Get("name")
  if name == ""{
    name = "Manoj"
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