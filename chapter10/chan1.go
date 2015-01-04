// channels allow two go routines to communicate
// with each other and synchronize their actions

package main

import(
  "fmt"
  "time"
)

func pinger(c chan string){
  for i := 0; ;i++{
    c <- "ping"
  }
}

func printer(c <-chan string){
  for {
    msg := <- c
    fmt.Println("PRINTER: ", msg)
    time.Sleep(time.Second * 1)
  }
}

func ponger(c chan string){
  for i := 0;;i++{
    c <- "pong"
  }
}

func main(){
  var c chan string = make(chan string)

  go pinger(c)
  go ponger(c)
  go printer(c)

  var input string
  fmt.Scanln(&input)
}
