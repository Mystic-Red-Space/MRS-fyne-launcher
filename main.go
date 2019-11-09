package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type launcher struct {
	functions map[string]func()
	input     map[string]*widget.Entry
	buttons   map[string]*widget.Button
	window    fyne.Window
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
	c.input = make(map[string]*widget.Entry)

	return c
}

func (c *launcher) addInput(text string, password bool) *widget.Entry {
	input := widget.NewEntry()
	input.Password = password
	input.SetPlaceHolder(text)

	c.input[text] = input
	return input
}
func (c *launcher) addButton(text string, action func()) *widget.Button {
	button := widget.NewButton(text, action)
	c.buttons[text] = button

	return button
}

func (c *launcher) loadUI(app fyne.App) {
	var modpacks []modpack
	modpacknames := make([]string, 0, 5)
	getJson("https://api.mysticrs.tk/list", &modpacks)
	for _, modp := range modpacks {
		modpacknames = append(modpacknames, modp.Name)
	}
	c.window = app.NewWindow("MRS Launcher")
	c.window.SetIcon(resourceIconPng)
	c.window.SetContent(fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewVBox(
		c.addInput("E-mail", false),
		c.addInput("Password", true),
		c.addButton("Login", func() {

		}))))
	c.window.ShowAndRun()
}
func main() {
	application := app.New()
	c := newLauncher()
	c.loadUI(application)
}
