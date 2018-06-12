// Why do we use packages?

	//We use packages for a number of reasons. It helps with naming collisions, prepaering neat code, and custom packages can limit the resources needed for a program (e.g., one math function vs the whole math library)

// What is the difference between an identifier that starts with a capital letter and one which doesn’t? (Average vs average)

	//A capital letter is visible outside of the package. Identifiers with lower case letters are meant to not be seen outside and are more limited in scope.

// What is a package alias? How do you make one?
	
	//A package alias allows you to import a package of name 'x' instead as name 'y'. Just specify like: import y "gobook/chapter11/x"

// We copied the average function from chapter 7 to our new package. Create Min and Max functions which find the minimum and maximum values in a slice of float64s.

	//ok. see file:
	// func Min(xs []float64) float64 {
	// 	low := xs[0]
	// 	for _,x := range xs {
	// 	  if ( low > x) {
	// 		low = x
	// 	  }
	// 	}
	// 	return low
	//   }
	  
	//   func Max(xs []float64) float64 {
	// 	high := xs[0]
	// 	for _,x := range xs {
	// 	  if ( high < x) {
	// 		high = x
	// 	  }
	// 	}
	// 	return high
	//   }

// How would you document the functions you created in #3?
		//by adding a commit line  before the function inside math.go. viewable with godoc or heading to http://localhost:6060/pkg/gobook/chapter11/math/