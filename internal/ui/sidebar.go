package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type SideBar struct {
	router          *UiRouter
	sidebar         *fyne.Container
	buttons         []*widget.Button
	activeSplit_idx int
}

// this throws a compilation error if the struct doesnt follow the interface
var _ Component = (*SideBar)(nil)

func NewSideBar(r *UiRouter) *SideBar {
	buttonlist := make([]*widget.Button, 0)
	buttonlist = append(buttonlist, widget.NewButton("add split", func() {
		r.ShowForm()
	}))
	container := container.NewBorder(buttonlist[0], nil, nil, nil, layout.NewSpacer())
	return &SideBar{router: r, sidebar: container, buttons: buttonlist, activeSplit_idx: -1}
}

func (sb *SideBar) GetCanvas() fyne.CanvasObject {
	return sb.sidebar
}
