package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func main()  {
  res,_ := http.Get("http://localhost:5000")
  b,_ := ioutil.ReadAll(res.Body)
  defer res.Body.Close()
  fmt.Printf("%s\n",b)
}
