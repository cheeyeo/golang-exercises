package main

import(
  "fmt"
  "sort"
)

func greatest(x ...int) int{
  sort.Ints(x)
  return x[len(x)-1]
}

func main(){
  x := []int{300, 3, 2, 1000}
  fmt.Println(greatest(x...))
}
