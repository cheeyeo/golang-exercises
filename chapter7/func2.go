package main

import "fmt"

func average(xs []float64) float64{
  // panic("Not implemented!")
  total := 0.0
  for _, v := range xs{
    total += v
  }

  return total / float64(len(xs))
}

func main(){
  xs := []float64{98,93,77,82,83}
  fmt.Println(average(xs))
}
