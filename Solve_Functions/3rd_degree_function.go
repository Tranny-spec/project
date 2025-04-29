package solve_functions

import "math"

// func SolveThirdEquation(a,b,c,d float32) (int,[]float32) {
// 	if a==0 {
// 		return SolveSecondEquation(b,c,d)
// 	} else {
// 		var (
// 			solutions int
// 			answers   []float32
// 		)

//		}
//	}
func SolveThirdEquation(a, b, c, d float32) (int, []float32) {
	// Create a slice to hold the answers
	var answers []float32

	// Check if it's actually a lower degree equation
	if a == 0 {
		return SolveSecondEquation(b, c, d)
	}

	// Convert to depressed cubic t³ + pt + q = 0 via substitution x = t - b/(3a)
	// This simplifies the cubic formula significantly

	// Normalize coefficients
	b /= a
	c /= a
	d /= a

	// Calculate intermediary values
	p := (3.0*c - b*b) / 3.0
	q := (2.0*b*b*b - 9.0*b*c + 27.0*d) / 27.0
	offset := b / 3.0

	// Calculate discriminant
	discriminant := q*q/4.0 + p*p*p/27.0

	// Round to prevent floating point errors
	discriminant = float32(math.Floor(float64(discriminant)*1000000) / 1000000)

	// Case 1: One real root and two complex conjugate roots
	if discriminant > 0 {
		// Calculate cubic roots
		u := float32(-q/2.0 + math.Sqrt(float64(discriminant)))
		v := float32(-q/2.0 - math.Sqrt(float64(discriminant)))

		// Calculate cube roots, preserving signs
		u = float32(math.Cbrt(float64(u)))
		v = float32(math.Cbrt(float64(v)))

		// The real root is u + v - offset
		root := u + v - offset

		// Round to 3 decimal places
		root = float32(math.Floor(float64(root)*1000) / 1000)

		answers = append(answers, root)
		return 1, answers
	}

	// Case 2: All roots are real, at least two are equal (discriminant = 0)
	if discriminant == 0 {
		// Calculate the real roots
		if p == 0 {
			// Triple root
			root := -offset
			root = float32(math.Floor(float64(root)*1000) / 1000)
			answers = append(answers, root)
			return 1, answers
		}

		// One single root and one double root
		r1 := float32(2.0 * math.Cbrt(float64(q/2.0)))
		r2 := float32(-1.0 * math.Cbrt(float64(q/2.0)))

		// Apply the offset and round
		r1 = float32(math.Floor(float64(r1-offset)*1000) / 1000)
		r2 = float32(math.Floor(float64(r2-offset)*1000) / 1000)

		// Add roots to answer list
		if r1 < r2 {
			answers = append(answers, r1, r2)
		} else {
			answers = append(answers, r2, r1)
		}

		return 2, answers
	}

	// Case 3: All roots are real and distinct (discriminant < 0)
	// Use trigonometric solution for three real roots
	rho := math.Sqrt(float64(-p * p * p / 27.0))
	phi := math.Acos(float64(-q / (2.0 * float32(rho))))

	// The three roots
	r1 := float32(2.0 * math.Cbrt(rho) * math.Cos(phi/3.0))
	r2 := float32(2.0 * math.Cbrt(rho) * math.Cos((phi+2.0*math.Pi)/3.0))
	r3 := float32(2.0 * math.Cbrt(rho) * math.Cos((phi+4.0*math.Pi)/3.0))

	// Apply the offset and round
	r1 = float32(math.Floor(float64(r1-offset)*1000) / 1000)
	r2 = float32(math.Floor(float64(r2-offset)*1000) / 1000)
	r3 = float32(math.Floor(float64(r3-offset)*1000) / 1000)

	// Sort the roots
	answers = append(answers, r1, r2, r3)
	sortFloatSlice(answers)

	return 3, answers
}
