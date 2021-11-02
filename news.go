package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"io"
	"net/http"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var articleNum int = 1

func showNewsApp() {

	// API Work
	res, err := http.Get("https://newsapi.org/v2/top-headlines?sources=techcrunch&apiKey=954f14b25d7a4a5b9c3e2ba2a7bfc4c7")

	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	news, err := UnmarshalNews(body)
	if err != nil {
		fmt.Println(err)
	}

	img := canvas.NewImageFromFile("logo.jpg")
	img.FillMode = canvas.ImageFillOriginal

	label1 := canvas.NewText("News App", color.White)
	label1.TextStyle = fyne.TextStyle{Bold: true}
	label1.Alignment = fyne.TextAlignCenter
	label2 := widget.NewLabel(fmt.Sprintf("No of articles:%d",
		news.TotalResults))
	//show title
	label3 := widget.NewLabel(fmt.Sprintf("%s", news.Articles[1].Title))
	label3.TextStyle = fyne.TextStyle{Bold: true}
	label3.Wrapping = fyne.TextWrapBreak
	// show articles
	entry1 := widget.NewLabel(fmt.Sprintf("%s",
		news.Articles[1].Description))
	entry1.Wrapping = fyne.TextWrapBreak

	btn := widget.NewButton("Next", func() {
		articleNum += 1
		label3.Text = news.Articles[articleNum].Title
		entry1.Text = news.Articles[articleNum].Description
		label3.Refresh()
		entry1.Refresh()
	})

	newsApp := container.NewVBox(
		container.NewBorder(img, label2, nil, nil, container.NewVBox(
			label1,
			label3,
			entry1,
			btn,
		),
		),
	)
	w := myApp.NewWindow("News App")
	w.Resize(fyne.NewSize(500, 600))
	myApp.Settings().SetTheme(theme.DarkTheme())
	w.SetContent(
		container.NewBorder(nil, nil, nil, nil, newsApp),
	)
	w.Show()
}

func UnmarshalNews(data []byte) (News, error) {
	var r News
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *News) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type News struct {
	Status       string    `json:"status"`
	TotalResults int64     `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

type Article struct {
	Source      Source `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	URLToImage  string `json:"urlToImage"`
	PublishedAt string `json:"publishedAt"`
	Content     string `json:"content"`
}

type Source struct {
	ID   ID   `json:"id"`
	Name Name `json:"name"`
}

type ID string

const (
	Techcrunch ID = "techcrunch"
)

type Name string

const (
	TechCrunch Name = "TechCrunch"
)
