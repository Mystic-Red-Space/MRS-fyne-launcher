package main

import (
    "encoding/json"
    "fyne.io/fyne"
    "fyne.io/fyne/app"
    "fyne.io/fyne/layout"
    "fyne.io/fyne/theme"
    "fyne.io/fyne/widget"
    "io/ioutil"
    "net/http"
)

type launcher struct {
    functions map[string]func()
    modpacks  *widget.Select
    output    *widget.Label
    buttons   map[string]*widget.Button
    window    fyne.Window
}
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

func newLauncher() *launcher {
    c := &launcher{}
    c.functions = make(map[string]func())
    c.buttons = make(map[string]*widget.Button)

    return c
}

func (c *launcher) addButton(text string, action func()) *widget.Button {
    button := widget.NewButton(text, action)
    c.buttons[text] = button

    return button
}

func (c *launcher) loadUI(app fyne.App) {
    c.output = widget.NewLabel("Hello World!")
    resp, _ := http.Get("https://api.mysticrs.tk/list")
    var modpacks []modpack
    modpacknames := make([]string, 0, 5)
    res, _ := ioutil.ReadAll(resp.Body)
    _ = json.Unmarshal(res, &modpacks)
    for _, modp := range modpacks {
        modpacknames = append(modpacknames, modp.Name)
    }
    c.modpacks = widget.NewSelect(modpacknames, func(s string) {
        c.output.SetText(s)
    })
    c.window = app.NewWindow("MRS Launcher")
    c.window.SetIcon(resourceIconPng)
    c.window.SetContent(fyne.NewContainerWithLayout(layout.NewHBoxLayout(), fyne.NewContainerWithLayout(layout.NewVBoxLayout(), c.addButton("Home", func() {
        c.window.SetFullScreen(!c.window.FullScreen())
    }),
        c.addButton("Modpacks", func() {
            app.Settings().SetTheme(theme.LightTheme())
        }),
        c.addButton("Settings", func() {
            app.Settings().SetTheme(theme.DarkTheme())
        })), fyne.NewContainerWithLayout(layout.NewVBoxLayout(), widget.NewVBox(c.output, c.modpacks))))
    c.modpacks.SetSelected(modpacknames[0])
    c.window.ShowAndRun()
}
func main() {
    application := app.New()
    c := newLauncher()
    c.loadUI(application)
}
