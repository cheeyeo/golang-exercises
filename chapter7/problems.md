## Problems

1. sum is a function which takes a slice of numbers and adds them together. What would its function signature look like in Go?

```
sum(arr ...int) int
```

2. Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).

```
package main

import "fmt"

func half(x int)(int, bool){
  halved := x / 2

  if x % 2 == 0{
    return halved, true
  }else{
    return halved, false
  }
}

func main(){
  fmt.Println(half(1))
}
```

3. Write a function with one variadic parameter that finds the greatest number in a list of numbers.

```
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

```

4. Using makeEvenGenerator as an example,write a makeOddGenerator function that generates odd numbers.

```
package main

import "fmt"

func makeOddGenerator() func() uint{
  i := uint(1)
  return func()(ret uint){
    ret = i
    i += 2
    return
  }
}

func main(){
  nextOdd := makeOddGenerator()
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
  fmt.Println(nextOdd())
}

```

5. The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).

```
package main

import "fmt"

func fib(x uint) uint{
  if x == 0{
    return 0
  } else if x == 1 {
    return 1
  } else {
    return fib(x-1) + fib(x-2)
  }
}


func main(){
  fmt.Println(fib(6))
}
```

6. What are defer, panic and recover? How do you recover from a run-time panic?

defer schedules a function call to be run after a function is completed

panic causes a run time error which mainly indicates a programming issue much
like an exception

recover is to rescue the error caused by panic

since panic stops all execution, recover calls after it will not be run hence
we need to pair it up with defer

using all three we can construct a panic and recover block almost akin to a try/rescue
block:

```
package main
import "fmt"
func main() {
     defer func() {
           str := recover()
           fmt.Println(str)
     }()
     panic("PANIC")
}

```



