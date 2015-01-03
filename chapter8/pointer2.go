package main

import "fmt"

// use a pointer to change the actual passed in variable
func zero(x *int){
  *x = 0
}

func main(){
  x := 5
  zero(&x)
  fmt.Println(x)
}
