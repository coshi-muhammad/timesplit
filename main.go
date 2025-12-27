package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("test window")
	w.SetContent(widget.NewLabel("hello from go lang"))
	w.ShowAndRun()
}
