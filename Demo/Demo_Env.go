package main

import (
	"fmt"
	//Import functions to solve equations from Solve_Functions
	GUI "github.com/Tranny-spec/project/GUI"
)

func PrintResult(solutions int, answerslist []float32) {
	fmt.Println("Number of solutions: ", solutions)

	for i := 0; i < solutions; i++ {
		fmt.Println("Answer ", i+1, ": ", answerslist[i])
	}
}
func main() {
	// fmt.Println("Itegrated Development Enviroment for Assignment")
	// var solutions int
	// var answerslist []float32

	// solutions, answerslist = solveFunctions.SolveQuadraricEquation(2, -4, 5)
	// PrintResult(solutions, answerslist)

	GUI.DrawCubicGraph(2, -3, 5,-2, "Cubic Graph.png")

}

