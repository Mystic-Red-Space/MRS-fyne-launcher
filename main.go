package main

import (
    "fmt"
    "encoding/json"
    "fyne.io/fyne/app"
    "fyne.io/fyne/widget"
    "io/ioutil"
    "net/http"
)

type mcfile struct {
    Name string `json:"name"`
    Dir string `json:"dir"`
    Url string `json:"url"`
    Md5 string `json:"md5"`
}

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
    resp, _ := http.Get("https://api.mysticrs.tk/list")
    var modpacks []modpack
    modpacknames := make([]string, 0, 5)
    res, _ := ioutil.ReadAll(resp.Body)
    _ = json.Unmarshal(res, &modpacks)
    for _, modp := range modpacks {
	    modpacknames = append(modpacknames, modp.Name)
	}
    application := app.New()
    w := application.NewWindow("MRS Launcher")
    w.SetContent(widget.NewHBox(
        widget.NewVBox(
        widget.NewSelect(modpacknames, func(item string) {
            fmt.Println(item)
        }),
        widget.NewButton("Test", func() {
            if w.FullScreen() {
                w.SetFullScreen(false)
            } else {
                w.SetFullScreen(true)
            }
        })),
        widget.NewLabel("Hello world"),
        widget.NewButton("Quit", func() {
            application.Quit()
        }),
    ))

    w.ShowAndRun()
}
