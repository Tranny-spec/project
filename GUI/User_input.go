package GUI

// Title: Fyne GUI Toolkit
// From GitHub Repository – fyne-io/fyne
// URL: https://github.com/fyne-io/fyne
// License: BSD‑3‑Clause (“New” or “Revised” License)
// Copyright © 2018–2025 Fyne Authors
// Used under fair use for education provision

import (
	// Package for formatted I/O
	"fmt"
	// Package for conversions to and from string representations of basic data types
	"strconv"

	//Packages for GUI for application
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	//Import functions to solve equations from Solve_Functions
	solveFunctions "github.com/Tranny-spec/project/Solve_Functions"
)

// Function to format the printed results
func ShowAnswers(solutions int, answers []float32) string {
	// Declare an empty string to store the results.
	var results string

	// Check if there are no solutions.
	if solutions == 0 {
		results = "There is no root"

		// Check if there are complex roots.
	} else if solutions == 4 {
		results = "There are complex roots"

		// Check if there is only one root.
	} else if solutions == 1 {

		// Format the single root into the results string.
		results = "There is only one root; x= " + fmt.Sprintf("%g", answers[0])

		// If there are multiple roots.
	} else {
		results = "Roots: \n"

		// Iterate through each answer and concatenate to "results"
		for i := 0; i < solutions; i++ {
			results += fmt.Sprintf("x%d = %f", i+1, answers[i])

			// Add a semicolon if it's not the last root.
			if i != solutions-1 {
				results += "; "
			}
		}
	}

	// Return the formatted results string.
	return results
}

// Function to set up GUI for application
func ApplicationGUI() {
	// Create a Fyne application.
	appliation := app.New()

	// Create a window with a title.
	window := appliation.NewWindow("Equation calculator and Graph editor")

	// Resize the window
	window.Resize(fyne.NewSize(600, 400))

	// Create input fields for the first coefficient of the equation.
	inputFirstCoefficient := widget.NewEntry()
	inputFirstCoefficient.SetPlaceHolder("a")

	// Create input fields for the second coefficient of the equation.
	inputSecondCoefficient := widget.NewEntry()
	inputSecondCoefficient.SetPlaceHolder("b")

	// Create input fields for the third coefficient of the equation.
	inputThirdCoefficient := widget.NewEntry()
	inputThirdCoefficient.SetPlaceHolder("c")

	// Create input fields for the fourth coefficient of the equation.
	inputFourthCoefficient := widget.NewEntry()
	inputFourthCoefficient.SetPlaceHolder("d")

	// Create a menu to select the type of equation.
	equationType := widget.NewSelect([]string{"Linear", "Quadratic", "Cubic"}, func(string) {})

	//Set Linear as Default
	equationType.SetSelected("Linear")

	// Create a label to display results.
	output := widget.NewLabel("")

	// Create a label for Equation Graph
	graphImage := widget.NewLabel("Equation Graph: ")

	// Define input fields for each type of equation
	equationType.OnChanged = func(typeInput string) {

		// Show and hide coefficient input fields for each equation type

		if typeInput == "Linear" {
			inputThirdCoefficient.Hide()
			inputFourthCoefficient.Hide()

		} else if typeInput == "Quadratic" {
			inputThirdCoefficient.Show()
			inputFourthCoefficient.Hide()

		} else {
			inputThirdCoefficient.Show()
			inputFourthCoefficient.Show()
		}

	}

	// Create a vertical box container to select equation type and input coefficients
	form := container.NewVBox(
		widget.NewLabel("Select equation type:"),
		equationType,
		widget.NewLabel("Enter coefficients:"),
		inputFirstCoefficient, inputSecondCoefficient, inputThirdCoefficient, inputFourthCoefficient,
	)

	// Create a button for solving the equation and plotting the graph.
	solveButton := widget.NewButton("Solve and Plot", func() {

		// Format the input coefficients from text to float64.
		a, _ := strconv.ParseFloat(inputFirstCoefficient.Text, 64)
		b, _ := strconv.ParseFloat(inputSecondCoefficient.Text, 64)
		c, _ := strconv.ParseFloat(inputThirdCoefficient.Text, 64)
		d, _ := strconv.ParseFloat(inputFourthCoefficient.Text, 64)

		// Set the filename for the graph image.
		graphImageFile := "graphImageFile.png"

		// Declare a string to store the results.
		var results string

		// Processes based on the selected equation type.
		if equationType.Selected == "Linear" {

			// Draw the graph for a linear equation.
			DrawLinearGraph(a, b, graphImageFile)

			// Solve the linear equation.
			solutions, answers := solveFunctions.SolveLinearEquation(float32(a), float32(b))

			// Format the solutions for display
			results = ShowAnswers(solutions, answers)

		} else if equationType.Selected == "Quadratic" {

			// Draw the graph for a quadratic equation
			DrawQuadraticGraph(a, b, c, graphImageFile)

			// Solve the quadratic equation
			solutions, answers := solveFunctions.SolveQuadraricEquation(float32(a), float32(b), float32(c))

			// Format the solutions for display
			results = ShowAnswers(solutions, answers)
		} else if equationType.Selected == "Cubic" {

			// Draw the graph for a cubic equation.
			DrawCubicGraph(a, b, c, d, graphImageFile)

			// Solve the cubic equation.
			solutions, answers := solveFunctions.SolveCubicEquation(float32(a), float32(b), float32(c), float32(d))

			// Format the solutions for display.
			results = ShowAnswers(solutions, answers)
		}

		// Set the output label to the formatted results.
		output.SetText(results)

		// Create an image canvas from the generated graph file.
		image := canvas.NewImageFromFile(graphImageFile)

		// Set image fill mode.
		image.FillMode = canvas.ImageFillContain

		// Set minimum size for the image.
		image.SetMinSize(fyne.NewSize(600, 400))

		// Remove the old graph image placeholder and add the new image.
		form.Remove(graphImage)
		form.Add(image)

		// Refresh the form to update the UI.
		form.Refresh()
	})

	// Add the solve button, output label, and graph image placeholder to the form.
	form.Add(solveButton)
	form.Add(output)
	form.Add(graphImage)

	// Hide the input fields for the third and fourth coefficients in the beginning
	inputThirdCoefficient.Hide()
	inputFourthCoefficient.Hide()

	// Set the content of the window to the form.
	window.SetContent(form)

	// Show the window and run the application.
	window.ShowAndRun()
}
