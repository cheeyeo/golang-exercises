package main

import(
  "fmt"
  "sort"
)

func main(){
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }

  sort.Ints(x)

  fmt.Println(x)
  fmt.Println("Smallest int: ", x[0])
}
