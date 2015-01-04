// a go routine is a function
// that is capable of running concurrently with other functions
// use go keyword followed by function invocation

package main

import(
  "fmt"
  "time"
  "math/rand"
)

func f(n int){
  for i := 0; i < 10; i++{
    fmt.Println(n, ":", i)
    amt := time.Duration(rand.Intn(256))
    time.Sleep(time.Millisecond * amt)
  }
}

func main(){
  // run 10 go routines
  for i := 0; i < 10; i++{
    go f(i)
  }

  var input string
  fmt.Scanln(&input)
  fmt.Println("test: ", input)
}
