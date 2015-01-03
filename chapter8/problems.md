## Problems

1. How do you get the memory address of a variable?

We append an ampersand to the variable e.g. &x


2. How do you assign a value to a pointer?

We append asterik * to the variable e.g. *x = 1

3. How do you create a new pointer?

Use the new keyword

```go
x := new(int)
```

4. What is the value of x after running this program:

```go
func square(x *float64) {
  *x = *x * *x
}
func main() {
  x := 1.5
  square(&x)
}

```

x is 2.25

5. Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).

```go
package main

import "fmt"

func swap(x, y *int){
  *x, *y = *y, *x
}

func main(){
  x := 1
  y := 2

  fmt.Println(x,y)
  swap(&x, &y)
  fmt.Println(x,y)
}
```
