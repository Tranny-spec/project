package solve_functions

import "math"

//Definition of variable "solutions" in range 0 to 4,
//which is used to store number of solutions
//	0: No solution
//  1: One solution
//	2: Two solutions
//  3: Three solutions
//  4: Infinite solutions

// Solve Quadraric Equation
func SolveQuadraricEquation(a, b, c float32) (int, []float32) {
	// Declare "solutions" to store the number of solutions and
	// "answers" to store answers list
	var (
		solutions int
		answers   []float32
	)

	// Check if the a=0
	if a == 0 {
		// If a=0, the equation is linear (bx + c = 0).
		// Call SolveLinearEquation to handle this case.

		return SolveLinearEquation(b, c)
	} else {
		// Calculate the discriminant (delta = b^2 - 4ac).
		delta := b*b - 4*a*c

		// Determine the number of solutions based on the value of delta.
		if delta < 0 {
			// If delta is negative, there are no real solutions
			solutions = 0
		} else if delta == 0 {
			// If delta is zero, there is one unique solution
			solutions = 1

			// Calculate the single solution: x = -b / (2a)
			answer := -b / (2 * a)

			// Round the answer to three decimal places
			answers = append(answers, float32(math.Floor(float64(answer)*1000)/1000))
		} else {
			// If delta is positive, there are two distinct real solutions.
			solutions = 2

			// Calculate the square root of delta.
			sqrtDelta := float32(math.Sqrt(float64(delta)))

			// Calculate the two solutions
			// using the quadratic formula: x = (-b ± sqrt(delta)) / (2a).
			answer1 := (-b + sqrtDelta) / (2 * a)
			answer2 := (-b - sqrtDelta) / (2 * a)

			// Round answers to three decimal places.
			answer1 = float32(math.Floor(float64(answer1)*1000) / 1000)
			answer2 = float32(math.Floor(float64(answer2)*1000) / 1000)

			// Add the rounded answers to the answers list
			answers = append(answers, answer1, answer2)
		}
	}
	// Return the number of solutions and the answers list
	return solutions, answers
}
