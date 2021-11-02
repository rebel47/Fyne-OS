package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"strconv"

	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showWeatherApp(w fyne.Window) {
	// a := app.New()
	// w := a.NewWindow("Weather App")
	// w.Resize(fyne.NewSize(800, 400))

	// API PART

	img := canvas.NewImageFromFile("weather.png")
	img.FillMode = canvas.ImageFill(500)
	label1 := canvas.NewText("Weather Details", color.Black)
	label1.Alignment = fyne.TextAlignCenter
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label1.TextSize = 25
	label2 := canvas.NewText(fmt.Sprintf("Country %s", "--"), color.Black)
	label2.Alignment = fyne.TextAlignCenter
	label2.TextSize = 25
	label3 := canvas.NewText(fmt.Sprint("Last Updated  ", "--"), color.Black)
	label3.Alignment = fyne.TextAlignCenter
	label3.TextSize = 25
	label4 := canvas.NewText(fmt.Sprintf("Humidity   %s", "--"), color.Black)
	label4.Alignment = fyne.TextAlignCenter
	label4.TextSize = 25
	label5 := canvas.NewText(fmt.Sprint("Temp   ", "--"), color.Black)
	label5.Alignment = fyne.TextAlignCenter
	label5.TextSize = 25
	label6 := canvas.NewText(fmt.Sprintf("City   %s", "--"), color.Black)
	label6.Alignment = fyne.TextAlignCenter
	label6.TextSize = 25
	// Location data using input box

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter your city name")
	input.TextStyle = fyne.TextStyle{Bold: true}

	content := container.NewVBox(input, widget.NewButton("Submit", func() {

		res, err := http.Get(fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=e7a6fabfe4ae403280593925210208&q=%s&aqi=yes", input.Text))
		if err != nil {
			fmt.Println(err)
		}
		body, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()
		weather, err := UnmarshalWeather(body)
		if err != nil {
			fmt.Println(err)
		}
		label6.Text = "City : " + (weather.Location.Name)
		label6.Refresh()
		label2.Text = "Country : " + (weather.Location.Country)
		label2.Refresh()
		label3.Text = "Last Updated : " + (weather.Current.LastUpdated)
		label3.Refresh()
		s := strconv.FormatInt(weather.Current.Humidity, 10)
		label4.Text = "Humidity : " + s + "%"
		label4.Refresh()
		label5.Text = "Temp : " + covertToString(float64(weather.Current.TempF)) + " F"
		label5.Refresh()

	}))
	// lin1 := container.New(layout.NewVBoxLayout(), label1)
	line2 := container.New(layout.NewVBoxLayout(), content)
	line3 := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), label1, label6, label2, label5, label4, label3, layout.NewSpacer())
	weatherApp := container.New(layout.NewMaxLayout(),
		// lin1,
		img,
		line2,
		line3,
	)
	w.SetContent(container.NewBorder(DeskBtn, nil, nil, nil, weatherApp))
}

func covertToString(val float64) string {
	res := strconv.FormatFloat(val, 'f', 2, 64)
	return res
}

func UnmarshalWeather(data []byte) (Weather, error) {
	var r Weather
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Weather) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Weather struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Current struct {
	LastUpdatedEpoch int64              `json:"last_updated_epoch"`
	LastUpdated      string             `json:"last_updated"`
	TempC            int64              `json:"temp_c"`
	TempF            float64            `json:"temp_f"`
	IsDay            int64              `json:"is_day"`
	Condition        Condition          `json:"condition"`
	WindMph          int64              `json:"wind_mph"`
	WindKph          int64              `json:"wind_kph"`
	WindDegree       int64              `json:"wind_degree"`
	WindDir          string             `json:"wind_dir"`
	PressureMB       int64              `json:"pressure_mb"`
	PressureIn       float64            `json:"pressure_in"`
	PrecipMm         int64              `json:"precip_mm"`
	PrecipIn         int64              `json:"precip_in"`
	Humidity         int64              `json:"humidity"`
	Cloud            int64              `json:"cloud"`
	FeelslikeC       int64              `json:"feelslike_c"`
	FeelslikeF       float64            `json:"feelslike_f"`
	VisKM            float64            `json:"vis_km"`
	VisMiles         int64              `json:"vis_miles"`
	Uv               int64              `json:"uv"`
	GustMph          float64            `json:"gust_mph"`
	GustKph          float64            `json:"gust_kph"`
	AirQuality       map[string]float64 `json:"air_quality"`
}

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int64  `json:"code"`
}

type Location struct {
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	TzID           string  `json:"tz_id"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Localtime      string  `json:"localtime"`
}
