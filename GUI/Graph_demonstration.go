package GUI

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func drawGraph(title string, fn func(x float64) float64, outputFile string) {
	graph := plot.New()
	graph.Title.Text = title
	graph.X.Label.Text = "X"
	graph.Y.Label.Text = "Y"

	graph.X.Min = -20
	graph.X.Max = 20
	graph.Y.Min = -20
	graph.Y.Max = 20

	graph.X.Tick.Marker = plot.DefaultTicks{}
	graph.Y.Tick.Marker = plot.DefaultTicks{}

	pts := make(plotter.XYs, 100)
	for i := 0; i < 100; i++ {
		x := float64(i)/2 - 10
		pts[i].X = x
		pts[i].Y = fn(x)
	}

	line, _ := plotter.NewLine(pts)

	graph.Add(line)
	graph.Add(plotter.NewGrid())

	// Draw X axis at y=0
	xAxis := plotter.NewFunction(func(x float64) float64 { return 0 })
	xAxis.Color = plotter.DefaultLineStyle.Color
	xAxis.Width = vg.Points(1)
	graph.Add(xAxis)

	// Draw Y axis at x=0
	yAxisPts := plotter.XYs{
		{X: 0, Y: graph.Y.Min},
		{X: 0, Y: graph.Y.Max},
	}
	yAxis, _ := plotter.NewLine(yAxisPts)

	yAxis.Color = plotter.DefaultLineStyle.Color
	yAxis.Width = vg.Points(1)
	graph.Add(yAxis)

	graph.Save(8*vg.Inch, 8*vg.Inch, outputFile)

}

// DrawLinearGraph generates an image of a linear graph y = ax + b
func DrawLinearGraph(a, b float64, outputFile string) {
	drawGraph(
		fmt.Sprintf("Linear equation: %gx + %g", a, b),
		func(x float64) float64 { return a*x + b },
		outputFile,
	)
}

// DrawQuadraticGraph generates an image of a quadratic graph y = ax² + bx + c
func DrawQuadraticGraph(a, b, c float64, outputFile string) {
	drawGraph(
		fmt.Sprintf("Quadratic equation: %g*x^2 + %g*x + %g", a, b, c),
		func(x float64) float64 { return a*x*x + b*x + c },
		outputFile,
	)
}

func DrawCubicGraph(a, b, c, d float64, outputFile string) {
	drawGraph(
		fmt.Sprintf("Quadratic equation: %g*x^3+ %g*x^2 + %g*x+%g", a, b, c, d),
		func(x float64) float64 { return a*x*x*x + b*x*x + c*x + d },
		outputFile,
	)
}
