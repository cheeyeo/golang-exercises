package math

import "testing"

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

func TestAverage(t *testing.T) {
  for _, pair := range tests {
    v := Average(pair.values)
    if v != pair.average {
      t.Error(
        "For", pair.values,
        "expected", pair.average,
        "got", v,
      )
    }
  }
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

func TestAverageEmptyList(t *testing.T){
  v := Average([]float64{})
  if v != 0 {
    t.Error("Expected 0, got ", v)
  }
}

