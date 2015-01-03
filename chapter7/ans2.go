package main

import "fmt"

func half(x int)(int, bool){
  halved := x / 2

  if x % 2 == 0{
    return halved, true
  }else{
    return halved, false
  }
}

func main(){
  fmt.Println(half(1))
}

