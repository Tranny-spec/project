package solve_functions

import "math"

func SolveFirstEquation(a, b float32) (int, []float32) {
	var (
		solutions int
		answers   []float32
	)
	if a == 0 {
		if b != 0 {
			solutions = 0
			answers = append(answers, 4)
		} else {
			solutions = 4
			answers = append(answers, 0)
		}
	} else {
		answers = append(answers, float32(math.Floor(float64(-b/a)*1000)/1000))
	}

	return solutions, answers
}
