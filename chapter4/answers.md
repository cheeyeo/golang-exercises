###What are two ways to create a new variable?
```go
// these are all ways to create a new variable
var x int = 5
var y = 5
z := 5
```
###What is the value of x after running: 
```go
x := 5; x += 1?
```
x should be 6

###What is scope and how do you determine the scope of a variable in Go?
Scope is where you can use the variable... where it takes effect.
You can determine the scope by referencing the surrounding block of curly braces {}

###What is the difference between var and const?
Var is a standard variable declaration.
const is a constant, and cannot be changed

###Using the example program as a starting point, write a program that
###converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
```go
package main

import "fmt"

func main() {
  fmt.Print("Enter temp in Fahrenheit: ")
  var myinput float64
  fmt.Scanf("%f",&myinput)
  answer := ((myinput - 32)*5/9)
  fmt.Println(answer, " Celsius")
}
```

###Write another program that converts from feet into meters. (1 ft = 0.3048 m)
```go
package main

import "fmt"

func main() {
  fmt.Println("Enter number of feet:")
  var feet float64
  fmt.Scanf("%f",&feet)
  meters := feet * 0.3048
  fmt.Println(meters, " meters")
}
```
