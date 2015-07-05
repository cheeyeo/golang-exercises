###How are integers stored on a computer?
Integers are stored in binary, as zeroes and ones: 0101010011

####We know that (in base 10) the largest 1 digit number is 9 and the largest 2 digit number is 99. 
Given that in binary the largest 2 digit number is 11 (3), the largest 3 digit number is 111 (7) 
and the largest 4 digit number is 1111 (15) what's the largest 8 digit number? (hint: 101-1 = 9 and 102-1 = 99)
128+64+32+16+8+4+2+1

###Although overpowered for the task you can use Go as a calculator. Write a program that computes 32132 × 42452 and prints it to the terminal. (Use the * operator for multiplication)
```go
package main

import "fmt"

func main(){
  fmt.Println(32132 * 42452)
}
```

###What is a string? How do you find its length?
You can find a string's length using:
```go
len("this is a string")
// or:
var mystring string = "this is a string"
len(mystring)
```

###What's the value of the expression (true && false) || (false && true) || !(false && false)?
```go
(true && false) || (false && true) || !(false && false)
```
reduces to:
```go
(false) || (false) || !(false)
false || false || true
false || true
true
```
