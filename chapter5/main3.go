package main

import "fmt"

func main(){
  for i := 1; i <= 10; i++{
    if i % 2 == 0{ // gets the remainder
      fmt.Println(i, "even")
    }else{
      fmt.Println(i, "odd")
    }
  }
}
