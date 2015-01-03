## Problems

1. How do you access the 4th element of an array or slice?

```go

x := [5]int{ 1, 2, 3, 4, 5 }
x[3]

// for a slice
x[:3], x[0:3]
```

2. What is the length of a slice created using make([]int, 3, 9)?

Length 3


3. Given the array x := [6]string{"a","b","c","d","e","f"}, what would
x[2:5] give you?

["c", "d", "e", "f"]

4. Write a program that finds the smallest number in this list:
```
x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
}
```



