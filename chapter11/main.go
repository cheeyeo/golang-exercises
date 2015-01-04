// need to export the GOPATH
// export GOPATH=/Users/chee/golang/golang-book/chapter11
// math module needs to go into src folder

// to generate godoc:
// godoc src/chapter11/math Average
// godoc -http=":6060"
// http://localhost:6060/pkg

package main

import (
  "fmt"
  m "chapter11/math"
)


func main(){
  xs := []float64{1,2,3,4}
  avg := m.Average(xs)
  fmt.Println(avg)
  fmt.Println(m.Min(xs))
  fmt.Println(m.Max(xs))
}
