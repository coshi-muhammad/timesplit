package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/core"
)

type FormView struct {
	logic  core.Logic
	router *UiRouter
	label  *widget.Label
	entry  *widget.Entry
	form   *fyne.Container
}

// this throws a compilation error if the struct doesnt follow the interface
var _ View = (*FormView)(nil)

func NewFormView(l core.FormLogic, r *UiRouter) *FormView {
	label := widget.NewLabel("Enter your message to be saved")
	entry := widget.NewEntry()
	container := container.New(layout.NewFormLayout(), label, entry)
	fv := &FormView{
		logic:  l,
		router: r,
		label:  label,
		entry:  entry,
		form:   container,
	}
	return fv
}

func (fv *FormView) GetCanvas() fyne.CanvasObject {
	submit := widget.NewButton("Submit", func() {
		fv.logic.SetFields(fv.entry.Text)
		fv.logic.Submit()
		fv.logic.Log()
		fv.router.sidebar = NewSideBar(fv.router)
		fv.router.ShowEmpty()
	})
	view := container.NewVBox(fv.form, submit)
	return view
}
