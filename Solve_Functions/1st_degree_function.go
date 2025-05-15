package solveFunctions

import "math"

//Definition of variable "solutions" in range 0 to 4,
//which is used to store number of solutions
//	0: No solution
//  1: One solution
//	2: Two solutions
//  3: Three solutions
//  4: Infinite solutions

// Function to solve linear equation
func SolveLinearEquation(a, b float32) (int, []float32) {
	// Declare "solutions" to store the number of solutions and
	// "answers" to store answers list
	var (
		solutions int
		answers   []float32
	)

	// Check if a=0
	if a == 0 {
		// If a=0, the equation becomes 0*x + b = 0, which simplifies to b = 0.
		if b != 0 {
			// If b is not zero (e.g., 0*x + 1000 = 0),
			// there is no value of x that can satisfy the equation.
			solutions = 0
		} else {
			// If b=0 (e.g., 0*x + 0 = 0),
			// any value of x satisfies the equation.
			solutions = 4
		}
	} else {
		// If a is not zero, there is exactly one solution: x = -b / a.
		solutions = 1

		// Calculate the answer and round it to three decimal places..
		answers = append(answers, float32(math.Floor(float64(-b/a)*1000)/1000))
	}

	// Return the number of solutions and the answers list
	return solutions, answers
}
