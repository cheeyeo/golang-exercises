package main

import "fmt"

func main(){
  arr := [5]float64{
    98,
    93,
    77,
    82,
    83,
  }

  x := arr[:5]

  var total float64 = 0

  for _, value := range x{
    total += value
  }

  fmt.Println(total / float64(len(x)))
}
