## What is whitespace?

Characters such as newlines, spaces and tabs are known as whitespace
as it doesn't get displayed when the prog is run. It is used to indent
code to make it easier to read.

## What is a comment? What are the two main ways of writing comments in Go?

A comment is a human readable section of text which is not executed by the
compiler and is used to aid in making clear what sections of code do.

One way to write a comment is to use the

```go
'//'
```

Another way is to use

```go
'/* ... */'
```

which is used for multi line comments


## Our program begins with the package main. What would the files in the fmt package begin with?

The files would begin with
```go
'package fmt'.
```

## We used the Println function defined in the fmt package. if we wanted to use the Exit function from the os package, what would we need to do?

```go
import "os"

os.Exit(0)
```




