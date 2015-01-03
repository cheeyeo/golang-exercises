package main

import "fmt"

// use a pointer to change the actual passed in variable
func one(x *int){
  *x = 1
}

func main(){
  x := new(int)
  one(x)
  fmt.Println(*x)
}
