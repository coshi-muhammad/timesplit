package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/core"
)

var _ View = (*EmptyView)(nil)

type EmptyView struct {
	logic        core.EmptyLogic
	router       *UiRouter
	button       *widget.Button
	emptySection *fyne.Container
}

func NewEmptyView(l core.Logic, r *UiRouter) *EmptyView {
	button := widget.NewButton("add split", func() {
		popup := NewSplitDialauge(r)
		popup.Show()
	})
	container := container.NewStack(button)
	return &EmptyView{logic: l, router: r, button: button, emptySection: container}
}

func (ev *EmptyView) GetCanvas() fyne.CanvasObject {
	return ev.emptySection
}
