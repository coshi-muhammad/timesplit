package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/logic"
	"github.com/google/uuid"
)

type SideBar struct {
	router          *UiRouter
	sidebar         *fyne.Container
	buttons         *widget.List
	activeSplit_idx int
}

// this throws a compilation error if the struct doesnt follow the interface
var _ Component = (*SideBar)(nil)

func NewSideBar(r *UiRouter) *SideBar {
	addButton := widget.NewButton("Add Split", func() {
		popup := NewSplitDialauge(r)
		popup.Show()
	})
	keys := make([]uuid.UUID, 0, len(r.GetState().Split_List))
	for k := range r.GetState().Split_List {
		keys = append(keys, k)
	}
	buttonlist := widget.NewList(
		func() int {
			return len(r.GetState().Split_List)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("template", func() {

			})
		},
		func(lii widget.ListItemID, co fyne.CanvasObject) {
			key := keys[lii]
			split := r.GetState().Split_List[key]
			co.(*widget.Button).SetText(split.Title)
			co.(*widget.Button).OnTapped = func() {
				r.Logic.SetActiveSplit(split.Id)
				r.Logic.GetState().Config.Active_Split_id = split.Id
				daylogic := logic.NewDayService(r.Logic,
					logic.WrapSplit(r.Logic.GetActiveSplit()))
				r.ShowDay(daylogic)
			}
		},
	)
	container := container.NewBorder(addButton, nil, nil, nil, buttonlist)
	return &SideBar{router: r, sidebar: container, buttons: buttonlist, activeSplit_idx: -1}
}

func (sb *SideBar) GetCanvas() fyne.CanvasObject {
	return sb.sidebar
}
