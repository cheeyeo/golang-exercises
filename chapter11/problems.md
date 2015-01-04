## problems

1. Why do we use packages?

* Namespace modules
* Code reuse

2. What is the difference between an identifier that starts with a capital letter and one which doesn’t? (Average vs average)

the lowercase function means that the function is private to the scope
of the defined package whereas an uppercase functions makes it public
to be used outside of the package

3. What is a package alias? How do you make one?

package alias prevents conflict with existing system packages
e.g. math package versus self defined maths package

using the above example, if we have created our own maths package, we can create
a package alias like so:

```go
import mymath "mymathpackagepath"
```

4. We copied the average function from chapter 7 to our new package.
Create Min and Max functions which find the minimum and maximum values in a slice of float64s.

```go
import "sort"

func Min(xs []float64) float64{
  sort.Float64s(xs)
  return xs[0]
}

func Max(xs []float64) float64{
  sort.Float64s(xs)
  return xs[len(xs)-1]
}
```


5. How would you document the functions you created in #3?

Create comments before the functions

```go
godoc golang-book/chapter11/math

//OR

godoc golang-book/chapter11/math Average Min Max
```

