## Problems

1. What's the difference between a method and a function?

A method is linked to a struct and can only be invoked on a struct
whereas a function can be invoked from anywhere it is declared

2. Why would you use an embedded anonymous field instead of a normal named field?

An embedded anonymous field reflects a is-a relationship intuitively

3. Add a new method to the Shape interface called perimeter which calculates the perimeter of a shape. Implement the method for Circle and Rectangle.

```go
type Shape interface{
  area() float64
  perimeter() float64
}

func (c *Circle) perimeter() float64{
  return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64{
  l := distance(r.x1, r.y1, r.x1, r.y2)
  w := distance(r.x1, r.y1, r.x2, r.y1)
  return (2 * l) + (2 * w)
}
```
