package main

import "fmt"
<<<<<<< HEAD
import "golang-book/chapter11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
=======
import "gobook/chapter11/math"

func main() {
  xs := []float64{1,2,3,4,600}

  avg := math.Average(xs)
  fmt.Println(avg)  
  
  low := math.Min(xs)
  fmt.Println(low)  
  
  high := math.Max(xs)
  fmt.Println(high)
}
>>>>>>> ca270e7c87ee5e17817083acf43605886265c6ff
