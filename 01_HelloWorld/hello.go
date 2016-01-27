package main

import (
  "fmt"
)

func main() {
  name := ""
  fmt.Printf("What's your name? ")
  fmt.Scanf("%s", &name)
  fmt.Printf("Hello, %v, nice to meet you!\n", name)
}
