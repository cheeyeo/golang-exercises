## Problems

1. Writing a good suite of tests is not always easy, but the process of writings tests often reveals more about a problem then you may at first realize.
For example, with our Average function what happens if you pass in an empty list ([]float64{})? How could we modify the function to return 0 in this case?


```go
// math.go

func Average(xs []float64) float64{
  if len(xs) < 1{
    return 0
  }

  total := float64(0)
  for _, x := range xs{
    total += x
  }

  return total / float64(len(xs))
}


func TestAverageEmptyList(t *testing.T){
  v := Average([]float64{})
  if v != 0 {
    t.Error("Expected 0, got ", v)
  }
}

```

2. Write a series of tests for the Min and Max functions you wrote in the previous chapter.

```go
type testpair struct{
  values []float64
  average float64
  min float64
  max float64
}

var tests = []testpair{
  { []float64{1,2}, 1.5, 1, 2},
  { []float64{1,1,1,1,1,1}, 1, 1, 1},
  { []float64{-1,1}, 0, -1, 1},
}

func TestMin(t *testing.T) {
  for _, pair := range tests {
    v := Min(pair.values)
    if v != pair.min {
      t.Error(
        "For", pair.values,
        "expected", pair.min,
        "got", v,
      )
    }
  }
}

func TestMax(t *testing.T) {
  for _, pair := range tests {
    v := Max(pair.values)
    if v != pair.max {
      t.Error(
        "For", pair.values,
        "expected", pair.max,
        "got", v,
      )
    }
  }
}
```
