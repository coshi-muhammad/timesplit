package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/coshi-muhammd/timesplit/internal/logic"
	"github.com/coshi-muhammd/timesplit/internal/ui"
)

func main() {
	a := app.New()
	w := a.NewWindow("form")
	formService := logic.NewFormService()
	router := ui.NewRouter(formService, w)
	router.ShowForm()
	w.ShowAndRun()
}
