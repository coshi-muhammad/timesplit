package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/coshi-muhammd/timesplit/internal/core"
	"github.com/coshi-muhammd/timesplit/internal/logic"
)

type UiRouter struct {
	Logic    core.RouterLogic
	isMobile bool
	App      fyne.App
	Window   fyne.Window
	Sidebar  *SideBar
	View     View
}

func NewRouter(l core.RouterLogic, w fyne.Window) *UiRouter {
	r := &UiRouter{Logic: l, Window: w}
	r.Logic.LoadSplits()
	//TODO: change this to use the splitservice
	r.Sidebar = NewSideBar(r)
	return r
}

func (r UiRouter) GetState() *core.AppState {
	return r.Logic.GetState()
}

func (r *UiRouter) ShowForm(l *logic.FormService) {
	form_view := NewFormView(l, r)
	r.View = form_view
	page := container.NewBorder(nil, nil, r.Sidebar.GetCanvas(),
		r.View.GetCanvas(), form_view.split_widget)
	r.Window.SetContent(page)
}

func (r *UiRouter) ShowDay(l *logic.DayService) {
	day_view := NewDayView(l, r)
	r.View = day_view
	page := container.NewBorder(nil, nil, r.Sidebar.GetCanvas(),
		day_view.GetCanvas(), day_view.SplitWidget)
	r.Window.SetContent(page)
}

func (r *UiRouter) ShowEmpty() {
	r.View = NewEmptyView(r.Logic, r)
	background := *canvas.NewRectangle(color.RGBA{0, 0, 0, 0})
	background.CornerRadius = 12
	background.StrokeWidth = 3
	background.StrokeColor = color.RGBA{0x8E, 0xAA, 0xCD, 0xFF}
	page := container.NewBorder(nil, nil, r.Sidebar.GetCanvas(), nil, r.View.GetCanvas(), &background)
	r.Window.SetContent(page)
}
