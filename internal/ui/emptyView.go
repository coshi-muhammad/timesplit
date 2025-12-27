package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/core"
)

type EmptyView struct {
	logic        core.Logic
	router       *UiRouter
	button       *widget.Button
	emptySection *fyne.Container
}

func NewEmptyView(l core.Logic, r *UiRouter) *EmptyView {
	button := widget.NewButton("add split", func() {
		r.ShowForm()
	})
	container := container.NewStack(button)
	return &EmptyView{logic: l, router: r, button: button, emptySection: container}
}

func (ev *EmptyView) GetCanvas() fyne.CanvasObject {
	return ev.emptySection
}
