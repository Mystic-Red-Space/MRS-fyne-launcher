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
	w := application.NewWindow("MRS Launcher")
	w.SetContent(widget.NewHBox(
		widget.NewVBox(widget.NewButtonWithIcon("Test",icon, func() {
			if w.FullScreen() {
				w.SetFullScreen(false)
			} else {
				w.SetFullScreen(true)
			}
		})),
		widget.NewLabel("Hello Fyne!"),
		widget.NewButton("Quit", func() {
			application.Quit()
		}),
	))

	w.ShowAndRun()
}
