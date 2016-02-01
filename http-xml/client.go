package main

import (
  "net/http"
  "encoding/xml"
  "bytes"
  "fmt"
  "io/ioutil"
)

type Person struct {
  Name string
  Age int `xml:",attr"`
}

func main() {
  person := &Person{"John", 27}
  buf, _ := xml.Marshal(person)
  body := bytes.NewBuffer(buf)
  r, _ := http.Post("http://localhost:8080", "text/xml", body)
  response, _ := ioutil.ReadAll(r.Body)
  defer r.Body.Close()
  fmt.Println(string(response))
}
