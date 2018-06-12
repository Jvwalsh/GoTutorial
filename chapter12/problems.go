// Writing a good suite of tests is not always easy, but the process of writings tests often reveals more about a problem then you may at first realize. For example, with our Average function what happens if you pass in an empty list ([]float64{})? How could we modify the function to return 0 in this case?

	//if we pass {} to the average function, it will return NaN because we are trying to divide by zero. To correct this edge case, add a condition check at the start of the function average to determine if there is no length to out stuff len(xs) == 0 or so. Then return 0 if true. code correct in ./chapter11/

// Write a series of tests for the Min and Max functions you wrote in the previous chapter.

	//see code