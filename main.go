package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()

var myWindow fyne.Window = myApp.NewWindow("OS")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget
var btn5 fyne.Widget
var btn6 fyne.Widget
var btn7 fyne.Widget

var img fyne.CanvasObject

var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {
	myApp.Settings().SetTheme(theme.LightTheme())

	img = canvas.NewImageFromFile("2.jpg")

	// Weather App Icon
	weatherIcon, err := os.Open("Weathericon.png")
	if err != nil {
		log.Fatal(err)
	}
	weatherR := bufio.NewReader(weatherIcon)
	weatherB, err := ioutil.ReadAll(weatherR)
	if err != nil {
		log.Fatal(err)
	}

	btn1 = widget.NewButtonWithIcon("Weather App", fyne.NewStaticResource("weather icon", weatherB), func() {
		showWeatherApp(myWindow)
	})

	// Calculator App Icon
	calculatorIcon, err := os.Open("calcicon.png")
	if err != nil {
		log.Fatal(err)
	}
	calcR := bufio.NewReader(calculatorIcon)
	calcB, err := ioutil.ReadAll(calcR)
	if err != nil {
		log.Fatal(err)
	}

	btn2 = widget.NewButtonWithIcon("Calculator", fyne.NewStaticResource("calculator icon", calcB), func() {
		showCalc()
	})

	// Gallery App Icon
	galleryIcon, err := os.Open("galleryicon.png")
	if err != nil {
		log.Fatal(err)
	}
	galleryR := bufio.NewReader(galleryIcon)
	galleryB, err := ioutil.ReadAll(galleryR)
	if err != nil {
		log.Fatal(err)
	}

	btn3 = widget.NewButtonWithIcon("Gallery App", fyne.NewStaticResource("gallery icon", galleryB), func() {
		showGalleryApp(myWindow)
	})

	// Text Editor Icon
	texteditorIcon, err := os.Open("texticon.png")
	if err != nil {
		log.Fatal(err)
	}
	textR := bufio.NewReader(texteditorIcon)
	textB, err := ioutil.ReadAll(textR)
	if err != nil {
		log.Fatal(err)
	}
	btn4 = widget.NewButtonWithIcon("Text Editor", fyne.NewStaticResource("text editor icon", textB), func() {
		showTextEditor()
	})

	// News App Icon
	newsIcon, err := os.Open("newsicon.jpg")
	if err != nil {
		log.Fatal(err)
	}
	newsR := bufio.NewReader(newsIcon)
	newsB, err := ioutil.ReadAll(newsR)
	if err != nil {
		log.Fatal(err)
	}
	btn5 = widget.NewButtonWithIcon("News App", fyne.NewStaticResource("news icon", newsB), func() {
		showNewsApp()
	})

	// BMI App Icon
	bmiIcon, err := os.Open("bmiicon.png")
	if err != nil {
		log.Fatal(err)
	}
	bmiR := bufio.NewReader(bmiIcon)
	bmiB, err := ioutil.ReadAll(bmiR)
	if err != nil {
		log.Fatal(err)
	}

	btn6 = widget.NewButtonWithIcon("BMI Calculator", fyne.NewStaticResource("bmi icon", bmiB), func() {
		showBMICalc()
	})

	// Home Icon
	homeIcon, err := os.Open("home.png")
	if err != nil {
		log.Fatal(err)
	}
	homeR := bufio.NewReader(homeIcon)
	homeB, err := ioutil.ReadAll(homeR)
	if err != nil {
		log.Fatal(err)
	}
	DeskBtn = widget.NewButtonWithIcon("Home", fyne.NewStaticResource("home icon", homeB), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	// Quit Icon
	quitIcon, err := os.Open("quiticon.png")
	if err != nil {
		log.Fatal(err)
	}
	quitR := bufio.NewReader(quitIcon)
	quitB, err := ioutil.ReadAll(quitR)
	if err != nil {
		log.Fatal(err)
	}
	btn7 = widget.NewButtonWithIcon("Shut Down", fyne.NewStaticResource("quit icon", quitB), func() {
		myApp.Quit()
	})

	panelContent = container.NewVBox((container.NewGridWithColumns(8, DeskBtn, btn1, btn2, btn3, btn4, btn5, btn6, btn7)))

	myWindow.Resize(fyne.NewSize(1280, 720))
	// myWindow.SetFullScreen(true) // For full screen
	myWindow.CenterOnScreen()
	myWindow.SetContent(
		container.NewBorder(panelContent, nil, nil, nil, img),
	)
	myWindow.ShowAndRun()

}
