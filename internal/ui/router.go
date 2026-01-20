package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"github.com/coshi-muhammd/timesplit/internal/core"
)

type UiRouter struct {
	// TODO: maybe change latter so you dont modify directly and
	// only use the logic interface to talk to it
	state   core.AppState
	window  fyne.Window
	logic   core.Logic
	sidebar *SideBar
	view    View
}

func NewRouter(l core.Logic, w fyne.Window) *UiRouter {
	r := &UiRouter{logic: l, window: w}
	r.sidebar = NewSideBar(r)
	return r
}

func (r *UiRouter) ShowForm() {
	r.view = NewFormView(r.logic, r)
	page := container.NewHBox(r.sidebar.GetCanvas(), r.view.GetCanvas())
	r.window.SetContent(page)
}
func (r *UiRouter) ShowEmpty() {
	r.view = NewEmptyView(r.logic, r)
	page := container.NewHBox(r.sidebar.GetCanvas(), r.view.GetCanvas())
	r.window.SetContent(page)
}
