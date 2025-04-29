package solve_functions

import (
	"math"
)

func SolveThirdEquation(a, b, c, d float32) (int, []float32) {

	var (
		solutions int
		answers   []float32
	)

	if a == 0 {
		return SolveSecondEquation(b, c, d)
	} else {
		b /= a
		c /= a
		d /= a

		p := (3.0*c - b*b) / 3.0
		q := (2.0*b*b*b - 9.0*b*c + 27.0*d) / 27.0
		offset := b / 3.0

		delta := q*q/4.0 + p*p*p/27.0

		delta = float32(math.Floor(float64(delta)*1000000) / 1000000)

		if delta > 0 {
			u := float32(float64(-q/2.0) + math.Sqrt(float64(delta)))
			v := float32(float64(-q/2.0) - math.Sqrt(float64(delta)))

			u = float32(math.Cbrt(float64(u)))
			v = float32(math.Cbrt(float64(v)))

			root := u + v - offset

			root = float32(math.Floor(float64(root)*1000) / 1000)

			solutions = 1
			answers = append(answers, root)
		} else if delta == 0 {
			if p == 0 {
				root := -offset
				root = float32(math.Floor(float64(root)*1000) / 1000)

				solutions = 1
				answers = append(answers, root)

			}

			r1 := float32(2.0 * math.Cbrt(float64(q/2.0)))
			r2 := float32(-1.0 * math.Cbrt(float64(q/2.0)))

			r1 = float32(math.Floor(float64(r1-offset)*1000) / 1000)
			r2 = float32(math.Floor(float64(r2-offset)*1000) / 1000)

			solutions = 2

			if r1 < r2 {
				answers = append(answers, r1, r2)
			} else {
				answers = append(answers, r2, r1)
			}

		} else {
			rho := math.Sqrt(float64(-p * p * p / 27.0))
			phi := math.Acos(float64(-q / (2.0 * float32(rho))))

			r1 := float32(2.0 * math.Cbrt(rho) * math.Cos(phi/3.0))
			r2 := float32(2.0 * math.Cbrt(rho) * math.Cos((phi+2.0*math.Pi)/3.0))
			r3 := float32(2.0 * math.Cbrt(rho) * math.Cos((phi+4.0*math.Pi)/3.0))

			r1 = float32(math.Floor(float64(r1-offset)*1000) / 1000)
			r2 = float32(math.Floor(float64(r2-offset)*1000) / 1000)
			r3 = float32(math.Floor(float64(r3-offset)*1000) / 1000)

			solutions = 3
			answers = append(answers, r1, r2, r3)
		}
	}

	return solutions, answers
}
