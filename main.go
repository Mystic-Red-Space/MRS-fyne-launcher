package main

import (
	"encoding/json"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"io/ioutil"
	"net/http"
)

type modpack struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
		Profile string `json:"profile"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	resp, err := http.Get("https://api.mysticrs.tk/list")
	var modpacklist []modpack
	res, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(res, &modpacklist)
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
