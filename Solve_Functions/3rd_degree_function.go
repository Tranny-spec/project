package solve_functions

import (
	"math"
)

//Definition of variable "solutions" in range 0 to 4,
//which is used to store number of solutions
//	0: No solution
//  1: One solution
//	2: Two solutions
//  3: Three solutions
//  4: Infinite solutions

// Title: Cubic Equations
// From Brilliant Math & Science Wiki – Cubic Equations
// URL: https://brilliant.org/wiki/cubic-equations/ 
// Authors: Arkajyoti Banerjee, Jenna Nieminen, Ayush G Rai, Jimin Khim, and 1 other 
// Copyright © Brilliant.org. 
// Used under fair use for education provision 

// Function to solve cubic equation
func SolveCubicEquation(a, b, c, d float32) (int, []float32) {
	// Declare "solutions" to store the number of solutions and
	// "answers" to store answers list
	var (
		solutions int
		answers   []float32
	)

	// Check if a=0
	if a == 0 {
		// If a=0, the equation is quadratic (bx^2 + cx + d = 0)
		// Call SolveQuadraricEquation to handle this case
		return SolveQuadraricEquation(b, c, d)

	} else {
		// Normalise the equation by dividing all coefficients
		// by a to get x^3 + (b/a)x^2 + (c/a)x + (d/a) = 0.
		// The equation becomes x^3 + Bx^2 + Cx + D = 0
		b /= a // B
		c /= a // C
		d /= a // D

		// Transform into a "depressed cubic" form: y^3 + py + q = 0
		// by substituting x = y - B/3.
		// Calculate p and q
		p := (3.0*c - b*b) / 3.0
		q := (2.0*b*b*b - 9.0*b*c + 27.0*d) / 27.0

		// The offset is used to convert solutions of the depressed cubic
		// back to solutions of the original normalised cubic.
		offset := b / 3.0

		// Calculate the discriminant of the depressed cubic: delta = (q/2)^2 + (p/3)^3.
		delta := q*q/4.0 + p*p*p/27.0

		// Round delta to 6 decimal places
		delta = float32(math.Floor(float64(delta)*1000000) / 1000000)

		// if delta > 0, there is one real root and two complex conjugate roots.
		if delta > 0 {

			sqrtDelta := float32(math.Sqrt(float64(delta)))

			u := float32(math.Cbrt(float64(-q/2.0 + sqrtDelta)))
			v := float32(math.Cbrt(float64(-q/2.0 - sqrtDelta)))

			// Convert back to the root of the original equation: x = y - offset.
			root := u + v - offset

			// Round the root to three decimal places.
			root = float32(math.Floor(float64(root)*1000) / 1000)

			//Set solutions and append answers into answers list
			solutions = 1
			answers = append(answers, root)

			// if delta == 0, there is onoe root
		} else if delta == 0 {

			if p == 0 {
				// q is also 0 if delta is 0 and p is 0.
				// y^3 = 0, so y = 0. This is a triple root.
				// The root of the original equation is x = -offset.
				root := -offset
				root = float32(math.Floor(float64(root)*1000) / 1000)

				//Set solutions and append answers into answers list
				solutions = 1
				answers = append(answers, root)
				//if p!=0, there are two roots
			} else {

				//Calculate two roots
				depressedRoot1 := float32(2.0 * math.Cbrt(float64(q/2.0)))
				depressedRoot2 := float32(-1.0 * math.Cbrt(float64(q/2.0)))

				// Convert back to roots of the original equation and round.
				root1 := float32(math.Floor(float64(depressedRoot1-offset)*1000) / 1000)
				root2 := float32(math.Floor(float64(depressedRoot2-offset)*1000) / 1000)

				solutions = 2

				//Set solutions and append answers into answers list
				if root1 == root2 {
					answers = append(answers, root1, root2)
				} else {
					answers = append(answers, root1)
				}
			}
			// If delta < 0, there are three roots
		} else {

			//Caculate rho,  -p * p * p / 27.0
			rho := math.Sqrt(float64(-p * p * p / 27.0))

			// Argument for acos: (-q/2) / rho
			// Ensure it's within [-1, 1]
			acos_arg := float64(-q / (2.0 * float32(rho)))
			if acos_arg > 1.0 {
				acos_arg = 1.0
			}
			if acos_arg < -1.0 {
				acos_arg = -1.0
			}
			phi := math.Acos(acos_arg)

			// Calculate the three real roots of the depressed cubic.
			sqrtNegPDiv3 := float32(math.Sqrt(float64(-p / 3.0)))

			depressedRoot1 := 2.0 * sqrtNegPDiv3 * float32(math.Cos(phi/3.0))
			depressedRoot2 := 2.0 * sqrtNegPDiv3 * float32(math.Cos((phi+2.0*math.Pi)/3.0))
			depressedRoot3 := 2.0 * sqrtNegPDiv3 * float32(math.Cos((phi+4.0*math.Pi)/3.0))

			// Convert back to roots of the original equation and round.
			root1 := float32(math.Floor(float64(depressedRoot1-offset)*1000) / 1000)
			root2 := float32(math.Floor(float64(depressedRoot2-offset)*1000) / 1000)
			root3 := float32(math.Floor(float64(depressedRoot3-offset)*1000) / 1000)

			//Set solutions and append answers into answers list
			solutions = 3.
			answers = append(answers, root1, root2, root3)
		}
	}

	// Return the number of solutions and the answers list
	return solutions, answers
}
