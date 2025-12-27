package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	current_time := time.Now().Format("The current time is: 03:04:05")
	clock.SetText(current_time)
}
func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	clock := widget.NewLabel("the current time is")
	go func() {
		for range time.Tick(time.Second * 1) {
			updateTime(clock)
		}
	}()
	w.SetContent(clock)
	w.SetMaster()
	w.Show()
	a.Run()
}
