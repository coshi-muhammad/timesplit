package ui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/logic"
)

var _ View = (*DayView)(nil)

type DayView struct {
	Logic           *logic.DayService
	router          *UiRouter
	SplitWidget     *SplitWidget
	title           *widget.Label
	time_range      *widget.Label
	description     *widget.Label
	atachement_list *widget.List
	sectionInfo     *fyne.Container
}

func formatTarget(input string) string {
	if _, err := os.Stat(input); err == nil {
		abs, _ := filepath.Abs(input)
		return "file://" + filepath.ToSlash(abs)
	}

	if !strings.Contains(input, "://") {
		return "https://" + input
	}

	return input
}

func (dv *DayView) SetSectionInfo() {
	dv.Logic.SetVisibleSection()
	active := widget.NewButton("Make active split", func() {
		dv.Logic.SetActiveSplit()
		fmt.Printf("active section is %s", dv.router.Logic.GetActiveSplit().Title)
	})
	fmt.Println(dv.Logic.Visible_section)
	dv.title = widget.NewLabel(dv.Logic.Visible_section.Title)
	start, end, _ := dv.Logic.Visible_section.GetFormatedTimestamps()
	time_string := fmt.Sprintf("%s - %s", start, end)
	dv.time_range = widget.NewLabel(time_string)
	dv.description = &widget.Label{}
	if dv.Logic.Visible_section.Description != "" {
		dv.description.Text = dv.Logic.Visible_section.Description
	}
	atatchment_label := widget.NewLabel("Atachement list")
	//TODO: refactor this later to use the data binding api
	// dv.atachement_list = widget.NewList(
	// 	func() int {
	// 		return len(dv.Logic.Visible_section.Atachements)
	// 	},
	// 	func() fyne.CanvasObject {
	// 		return widget.NewButton("template", func() {})
	// 	},
	// 	func(i widget.ListItemID, o fyne.CanvasObject) {
	// 		atatchment := dv.Logic.Visible_section.Atachements[i]
	// 		atatchment_link := formatTarget(atatchment)
	// 		o.(*widget.Button).SetText(atatchment)
	// 		o.(*widget.Button).OnTapped = func() {
	// 			url, _ := url.Parse(atatchment_link)
	// 			dv.router.App.OpenURL(url)
	// 		}
	// 	},
	// )
	dv.atachement_list = &widget.List{}

	dv.sectionInfo = container.NewBorder(active, dv.atachement_list, nil, nil,
		container.NewVBox(dv.title, dv.time_range, dv.description, atatchment_label))
}

func NewDayView(l *logic.DayService, r *UiRouter) *DayView {
	dv := &DayView{
		Logic:  l,
		router: r,
		//TODO: eventually either add it to the original or make a modification
		//of it here but make it that when you press a section you at least get its
		//uuid or the deafult non value if there is not section there
		SplitWidget: NewSplitWidget(l.Visible_split),
	}
	dv.SetSectionInfo()
	return dv
}

func (dv *DayView) GetCanvas() fyne.CanvasObject {
	return dv.sectionInfo
}
func (dv *DayView) SetVisibleSection() {
	dv.Logic.SetVisibleSection()
	dv.title.Text = dv.Logic.Visible_section.Title
	dv.title.Refresh()
	start, end, _ := dv.Logic.Visible_section.GetFormatedTimestamps()
	time_string := fmt.Sprintf("%s - %s", start, end)
	dv.time_range.Text = time_string
	dv.time_range.Refresh()
	dv.description.Text = dv.Logic.Visible_section.Description
	dv.description.Refresh()
}
