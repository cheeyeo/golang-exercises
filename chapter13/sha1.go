package main

import (
  "fmt"
  "io"
  "crypto/sha1"
)

func main(){
  h := sha1.New()
  io.WriteString(h, "Test")
  io.WriteString(h, "Test 2")
  fmt.Printf("% x\n", h.Sum(nil))

  data := []byte("This page intentionally left blank.")
  fmt.Printf("% x\n", sha1.Sum(data))
}
