## problems

1. How do you specify the direction of a channel type?

We can specify it using the redirection operator:

```go
c chan <- string

```

2. Write your own Sleep function using time.After.

```go
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

```

3. What is a buffered channel? How would you create one with a capacity of 20?

A buffered channel is asynchronous; sending or receiving a message will not wait
unless the channel is already full.

```go
c := make(chan int, 20)
```
