package main

import (
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"

	// "fyne.io/fyne/v2/internal/widget"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("BMI Calculator")
	w.Resize(fyne.NewSize(400, 300))
	a.Settings().SetTheme(theme.DarkTheme())

	// label
	label := canvas.NewText("BMI Calculator", color.Black)
	label.Alignment = fyne.TextAlignCenter

	result := canvas.NewText("", color.Black)
	result.Alignment = fyne.TextAlignCenter
	result.TextSize = 20

	// input Height
	inputHeight := widget.NewEntry()
	inputHeight.SetPlaceHolder("Enter the Height in Centimeter (cm)")

	// input Weight

	inputWeight := widget.NewEntry()
	inputWeight.SetPlaceHolder("Enter the Weight in Kilogram (KG)")

	// Button to calculate

	btn := widget.NewButton("Calculate", func() {
		h, _ := strconv.ParseFloat(inputHeight.Text, 64)
		w, _ := strconv.ParseFloat(inputWeight.Text, 64)

		result.Text = bmiCalc(h/100, w)
		result.Refresh()
	})

	w.SetContent(
		container.NewVBox(
			inputHeight,
			inputWeight,
			btn,
			result,
		),
	)
	w.ShowAndRun()
}
func bmiCalc(height, weight float64) string {

	var BMI = weight / math.Pow(height, 2)
	if BMI <= 18.4 {
		return "You are Under Weight"
	} else if BMI > 18.4 && BMI <= 25 {
		return "You are Healthy"

	} else if BMI > 25 && BMI <= 30 {
		return "You are Over Weight"

	} else if BMI > 30 && BMI <= 35 {
		return "You are Severely Over Weight"

	} else if BMI > 35 && BMI <= 40 {
		return "You are Obese"

	} else {
		return "You are Severely Obese"
	}
}

func covertToString(BMI float64) {
	panic("unimplemented")
}
