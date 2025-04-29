package solve_functions

import "math"

func SolveSecondEquation(a, b, c float32) (int, []float32) {

	var (
		solutions int
		answers   []float32
	)
	if a == 0 {
		return SolveFirstEquation(b, c)
	} else {
		delta := b*b - 4*a*c

		if delta < 0 {
			solutions = 0
			answers = append(answers, 0)
		} else if delta == 0 {
			solutions = 1

			answer := -b / (2 * a)
			answers = append(answers, float32(math.Floor(float64(answer)*1000)/1000))
		} else {
			solutions = 2

			sqrtDelta := float32(math.Sqrt(float64(delta)))
			answer1 := (-b + sqrtDelta) / (2 * a)
			answer2 := (-b - sqrtDelta) / (2 * a)

			answer1 = float32(math.Floor(float64(answer1)*1000) / 1000)
			answer2 = float32(math.Floor(float64(answer2)*1000) / 1000)

			answers = append(answers, answer1, answer2)
		}
	}
	return solutions, answers
}
