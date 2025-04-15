package SecondDegreeEquation

import (
	"math"
)

func SolveSecondEquation(a, b float32) float32 {

	answer := float32(math.Floor(float64(b/a)*1000) / 1000)

	return answer
}
