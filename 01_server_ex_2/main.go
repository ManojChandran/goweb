package main

import (
  "net/http"
  "fmt"
  "os"
)

func main()  {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/shutdown", shutdown)
  http.ListenAndServe(":5000", nil)
}

func homePage(res http.ResponseWriter, req *http.Request) {
  if req.URL.Path != "/"{
    http.NotFound(res, req)
    return
  }
  fmt.Fprintf(res, "Hi there, How r u?")
}

func shutdown(req http.ResponseWriter, res *http.Request) {
  os.Exit(0)
}
