package math

// Finds the average of a series of numbers
func Average(xs []float64) float64 {
  total := float64(0)
  for _, x := range xs {
    total += x
  }
  return total / float64(len(xs))
}

func Min(xs []float64) float64 {
  low := xs[0]
  for _,x := range xs {
    if ( low > x) {
      low = x
    }
  }
  return low
}

func Max(xs []float64) float64 {
  high := xs[0]
  for _,x := range xs {
    if ( high < x) {
      high = x
    }
  }
  return high
}