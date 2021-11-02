package main

import (
	// "image/color"
	"io/ioutil"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var count int = 1

func showTextEditor() {

	content := container.NewVBox(
		container.NewHBox(
			// widget.NewLabel("Pep Text Editor"),
			widget.NewLabelWithStyle("TEXT EDITOR", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
	)

	content.Add(widget.NewButton("Add New File", func() {
		content.Add(widget.NewLabel("New File" + strconv.Itoa(count)))
		count++
	}))
	// input := widget.NewEntry() this will take input only in one line for multiline entry use the code below
	input := widget.NewMultiLineEntry()
	input.SetPlaceHolder("Enter Text here.......")
	input.Resize(fyne.NewSize(400, 400))

	// to save file

	saveBtn := widget.NewButton("Save Text File", func() {
		saveFileDiolog := dialog.NewFileSave(
			func(uc fyne.URIWriteCloser, _ error) {
				textData := []byte(input.Text)
				uc.Write(textData)
			}, myWindow)

		saveFileDiolog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
		saveFileDiolog.Show()
	})

	openBtn := widget.NewButton("Open Text File", func() {
		openFileDialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				ReadData, _ := ioutil.ReadAll(r)

				output := fyne.NewStaticResource("New File", ReadData)

				viewData := widget.NewMultiLineEntry()

				viewData.SetText(string(output.StaticContent))
				w := fyne.CurrentApp().NewWindow(
					string(output.StaticName))

				inSaveBtn := widget.NewButton("Save", func() {
					inSaveFileDiolog := dialog.NewFileSave(
						func(uc fyne.URIWriteCloser, _ error) {
							textData := []byte(viewData.Text)
							uc.Write(textData)
						}, w)

					inSaveFileDiolog.SetFileName("New File" + strconv.Itoa(count-1) + ".txt")
					inSaveFileDiolog.Show()
				})
				w.SetContent(container.NewVBox(viewData, inSaveBtn))

				w.Resize(fyne.NewSize(400, 400))

				w.Show()
			}, myWindow)

		openFileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".txt"}))
		openFileDialog.Show()
	})

	textEditor := container.NewVBox(
		content,
		input,

		container.NewHBox(
			saveBtn,
			openBtn,
		),
	)

	w := myApp.NewWindow("Text Editor")
	w.Resize(fyne.NewSize(500, 280))
	myApp.Settings().SetTheme(theme.DarkTheme())
	// r, _ := fyne.LoadResourceFromPath("weather.png")
	// myApp.SetIcon(r)
	w.SetContent(
		container.NewBorder(nil, nil, nil, nil, textEditor),
	)
	w.Show()

}
