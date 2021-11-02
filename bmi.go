package main

import (
	"image/color"
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func showBMICalc() {

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
		result.Color = color.White
		result.TextStyle = fyne.TextStyle{Bold: true}
		result.Refresh()
	})

	calcBMI := container.NewVBox(
		inputHeight,
		inputWeight,
		btn,
		result,
	)
	// a := app.New()
	w := myApp.NewWindow("BMI Calculator")
	w.Resize(fyne.NewSize(300, 200))
	myApp.Settings().SetTheme(theme.DarkTheme())
	w.SetContent(
		container.NewBorder(nil, nil, nil, nil, calcBMI),
	)
	w.Show()
}
func bmiCalc(height, weight float64) string {

	var BMI = weight / math.Pow(height, 2)
	if BMI <= 18.4 {
		return "Your BMI is : " + covertToString(BMI) + " and You are Under Weight"
	} else if BMI > 18.4 && BMI <= 25 {
		return "Your BMI is : " + covertToString(BMI) + " and You are Healthy"

	} else if BMI > 25 && BMI <= 30 {
		return "Your BMI is : " + covertToString(BMI) + " and You are Over Weight"

	} else if BMI > 30 && BMI <= 35 {
		return "Your BMI is : " + covertToString(BMI) + " and You are Severely Over Weight"

	} else if BMI > 35 && BMI <= 40 {
		return "Your BMI is : " + covertToString(BMI) + " and You are Obese"

	} else {
		return "Your BMI is : " + covertToString(BMI) + " and You are Severely Obese"
	}
}
