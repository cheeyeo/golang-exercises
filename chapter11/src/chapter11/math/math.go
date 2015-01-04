package math

import "sort"

// Finds the average of a series of numbers
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

// Finds the minimum of a series of numbers
func Min(xs []float64) float64{
  sort.Float64s(xs)
  return xs[0]
}

// Finds the maximum of a series of numbers
func Max(xs []float64) float64{
  sort.Float64s(xs)
  return xs[len(xs)-1]
}
