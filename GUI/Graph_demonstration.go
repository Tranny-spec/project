package GUI

// Title: Gonum Numerical Libraries
// From GitHub Repository – gonum/gonum
// URL: https://github.com/gonum/gonum
// License: BSD-3-Clause “New” or “Revised” License
// Copyright © 2013 The Gonum Authors. All rights reserved.
// Used under fair use for education provision

import (
	// Package for formatted I/O.
	"fmt"

	// Gonum plot package for creating and editing plots
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// DrawGraph is a helper function to create and save a plot
// based on given mathematical function.
func DrawGraph(title string, equationFunction func(x float64) float64, graphImage string) {

	// Create a new plot.
	graph := plot.New()

	// Set the title of the graph.
	graph.Title.Text = title

	// Set the X-axis label.
	graph.X.Label.Text = "X"

	// Set the Y-axis label.
	graph.Y.Label.Text = "Y"

	// Set the range for the X-axis.
	graph.X.Min = -20
	graph.X.Max = 20

	// Set the range for the Y-axis.
	graph.Y.Min = -20
	graph.Y.Max = 20

	// Set default tick marks for x and y axes.
	graph.X.Tick.Marker = plot.DefaultTicks{}
	graph.Y.Tick.Marker = plot.DefaultTicks{}

	// Create a plotter.XYs
	pts := make(plotter.XYs, 100)

	//Iterate to generate data points for the graph
	for i := 0; i < 100; i++ {
		x := float64(i)/2.0 - 25

		pts[i].X = x

		// Calculate the corresponding y using the provided equation function
		pts[i].Y = equationFunction(x)
	}

	// Create a new line plotter from the points
	line, _ := plotter.NewLine(pts)

	// Add the line to the graph
	graph.Add(line)

	// Add a grid to the graph
	graph.Add(plotter.NewGrid())

	// Draw X-axis line (y=0)
	xAxis := plotter.NewFunction(func(x float64) float64 { return 0 })

	// Use default color
	xAxis.Color = plotter.DefaultLineStyle.Color

	// Set line width
	xAxis.Width = vg.Points(1)
	graph.Add(xAxis)

	// Draw Y-axis line (x=0)
	yAxisPts := plotter.XYs{
		// Start point of Y-axis
		{X: 0, Y: graph.Y.Min},

		// End point of Y-axis
		{X: 0, Y: graph.Y.Max},
	}

	//Define points for the Y-axis line
	yAxis, _ := plotter.NewLine(yAxisPts)

	// Use default color
	yAxis.Color = plotter.DefaultLineStyle.Color

	// Set line width
	yAxis.Width = vg.Points(1)
	graph.Add(yAxis)

	// Save the graph to a file with specified dimensions
	graph.Save(8*vg.Inch, 8*vg.Inch, graphImage)
}

// DrawLinearGraph generates and saves an image of a linear graph
func DrawLinearGraph(a, b float64, graphImage string) {
	// Call DrawGraph with the linear equation details.
	DrawGraph(
		fmt.Sprintf("Linear equation: y = %gx + %g", a, b), // Graph title.
		func(x float64) float64 { return a*x + b },         // Linear equation function.
		graphImage, // Filename for the graph image.
	)
}

// DrawQuadraticGraph generates and saves an image of a quadratic graph y = ax² + bx + c.
// a: The coefficient of x².
// b: The coefficient of x.
// c: The constant term.
// graphImage: The filename to save the generated graph image.
func DrawQuadraticGraph(a, b, c float64, graphImage string) {
	// Call drawGraph with the quadratic equation details.
	DrawGraph(
		fmt.Sprintf("Quadratic equation: y = %g*x^2 + %g*x + %g", a, b, c), // Graph title.
		func(x float64) float64 { return a*x*x + b*x + c },                 // Quadratic equation function.
		graphImage, // Filename for the graph image.
	)
}

// DrawCubicGraph generates and saves an image of a cubic graph y = ax³ + bx² + cx + d.
// a: The coefficient of x³.
// b: The coefficient of x².
// c: The coefficient of x.
// d: The constant term.
// graphImage: The filename to save the generated graph image.
func DrawCubicGraph(a, b, c, d float64, graphImage string) {
	// Call DrawGraph with the cubic equation details.
	DrawGraph(
		fmt.Sprintf("Cubic equation: y = %g*x^3 + %g*x^2 + %g*x + %g", a, b, c, d), // Graph title.
		func(x float64) float64 { return a*x*x*x + b*x*x + c*x + d },               // Cubic equation function.
		graphImage, // Filename for the graph image.
	)
}
