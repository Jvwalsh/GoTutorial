package math

// Finds the average of a series of numbers
func VarNum2(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}

	low := xs[0]
	high := xs[0]
	varNum2 := float64(0)
	// median := float64(0)
	for _, x := range xs {

		if x > high {
			high = x
		}
		if x < low {
			low = x
		}
	}

	varNum2 = high - low
	// median = higher - range/2;
	return varNum2
}
