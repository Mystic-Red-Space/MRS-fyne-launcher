package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	application := app.New()
	image, err := ioutil.ReadFile("icon.png")
	check(err)
	icon := fyne.NewStaticResource("icon",image)
	application.SetIcon(icon)
	w := application.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			application.Quit()
		}),
	))

	w.ShowAndRun()
}
