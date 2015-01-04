package main

import(
  "fmt"
  "time"
)

func mysleep(n int){
  select {
  case <-time.After(time.Duration(n)*time.Second):
    fmt.Println("End of Sleep")
  }
}

func main(){
  fmt.Println("Start")
  mysleep(3)
  fmt.Println("End")

  var input string
  fmt.Scanln(&input)
}
