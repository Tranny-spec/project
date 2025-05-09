// package main

// import (
// 	"fmt"

// 	GUI "github.com/Tranny-spec/project/GUI"
// 	solve_functions "github.com/Tranny-spec/project/Solve_Functions"
// 	"fyne.io/fyne/v2/app"
// 	"fyne.io/fyne/v2/internal/driver/glfw"
// 	"fyne.io/fyne/v2/internal/driver/common"
// 	"fyne.io/fyne/v2/internal/painter/gl"
// 	"github.com/go-gl/gl/v2.1/gl"
// )

// func main() {
// 	fmt.Println(solve_functions.SolveFirstEquation(3, 7))
// 	fmt.Println(solve_functions.SolveSecondEquation(2, 5, -10))
// 	fmt.Println(solve_functions.SolveThirdEquation(2, 5, -10, -100))

// 	GUI.DrawLinearGraph(2, 3, "Graph.png")
// 	GUI.DrawQuadraticGraph(2, 3, 10, "Graph1.png")
// 	GUI.DrawCubicGraph(2,3,10,-5,"Graph2.png")

// 	GUI.RunEquationApp()
// }

// filepath: /home/tranny/Documents/Project/project/main.go
package main

import (
    "fmt"

    GUI "github.com/Tranny-spec/project/GUI"
    solve_functions "github.com/Tranny-spec/project/Solve_Functions"
)

func main() {
    fmt.Println(solve_functions.SolveFirstEquation(3, 7))
    fmt.Println(solve_functions.SolveSecondEquation(2, 5, -10))
    fmt.Println(solve_functions.SolveThirdEquation(2, 5, -10, -100))

    // GUI.DrawLinearGraph(2, 3, "Graph.png")
    // GUI.DrawQuadraticGraph(2, 3, 10, "Graph1.png")
    // GUI.DrawCubicGraph(2, 3, 10, -5, "Graph2.png")

    GUI.ApplicationGUI()
}
