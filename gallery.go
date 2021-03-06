package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func showGalleryApp(w fyne.Window) {

	root_src := "C:\\Users\\aayya\\Desktop\\Pep Coding\\images"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
		log.Fatal(err)
	}
	tab := container.NewAppTabs()
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			extension := strings.Split(file.Name(), ".")[1]
			if extension == "png" || extension == "jpg" || extension == "jpeg" {
				image := canvas.NewImageFromFile(root_src + "\\" + file.Name())
				image.FillMode = canvas.ImageFill(500)
				tab.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}

	tab.SetTabLocation(container.TabLocationBottom)

	w.SetContent(
		container.NewBorder(DeskBtn, nil, nil, nil, tab),
	)
	w.Show()
}
