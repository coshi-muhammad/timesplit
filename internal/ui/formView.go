package ui

import (
	"fmt"
	"image/color"
	"slices"
	"unicode"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/logic"
)

type FormView struct {
	logic               *logic.FormService
	router              *UiRouter
	form                *widget.Form
	split_widget        *SplitWidget
	color_dialauge      *dialog.ColorPickerDialog
	atachement_dialauge *dialog.FileDialog
	atachement_list     *widget.List
	items               map[string]*widget.FormItem
}
type timeEntry struct {
	widget.Entry
}

func newTimeEntry() *timeEntry {
	e := &timeEntry{}
	e.ExtendBaseWidget(e)
	e.PlaceHolder = "HH:MM:SS"
	return e
}

// TypedRune intercepts every character typed
func (e *timeEntry) TypedRune(r rune) {
	// 1. Only allow numbers
	if !unicode.IsDigit(r) {
		return
	}

	// 2. Limit length to 8 characters (HH:MM:SS)
	if len(e.Text) >= 8 {
		return
	}

	// 3. Auto-insert colons at positions 2 and 5
	if len(e.Text) == 2 || len(e.Text) == 5 {
		e.Entry.TypedRune(':')
	}

	e.Entry.TypedRune(r)
}

// this throws a compilation error if the struct doesnt follow the interface
var _ View = (*FormView)(nil)

func NewFormView(l *logic.FormService, r *UiRouter) *FormView {
	fv := &FormView{
		logic:        l,
		router:       r,
		split_widget: NewSplitWidget(l.Temp_split),
		items:        make(map[string]*widget.FormItem),
	}
	fv.NewForm()
	return fv
}

func (fv *FormView) NewForm() {
	fv.items["title"] = widget.NewFormItem("Title:", widget.NewEntry())
	fv.items["start"] = widget.NewFormItem("Start:", newTimeEntry())
	fv.items["end"] = widget.NewFormItem("End:", newTimeEntry())
	fv.items["description"] = widget.NewFormItem("Description:", widget.NewMultiLineEntry())
	fv.color_dialauge = dialog.NewColorPicker("Color Picker", "Choose a color for the section",
		func(c color.Color) {
			fv.logic.Temp_section.Color =
				color.RGBAModel.Convert(c).(color.RGBA)
		}, fv.router.Window)
	fv.items["color"] = widget.NewFormItem("Color:", widget.NewButton("Pick Color", func() {
		fv.color_dialauge.Show()
	}),
	)

	fv.atachement_dialauge = dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
		if err == nil {
			fv.logic.Temp_section.Atachements =
				append(fv.logic.Temp_section.Atachements, reader.URI().Path())
			fv.atachement_list.Refresh()
		}
	}, fv.router.Window)
	fv.items["Atachement"] = widget.NewFormItem("Atachements", widget.NewButton("Add Atachement", func() {
		fv.atachement_dialauge.Show()
	}),
	)
	fv.atachement_list = widget.NewList(
		func() int {
			return len(fv.logic.Temp_section.Atachements)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("template"),
				layout.NewSpacer(),
				widget.NewButton("Remove", func() {}),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			hbox := o.(*fyne.Container)
			lbl := hbox.Objects[0].(*widget.Label)
			btn := hbox.Objects[2].(*widget.Button)

			lbl.SetText(fv.logic.Temp_section.Atachements[i])

			btn.OnTapped = func() {
				fv.logic.Temp_section.Atachements =
					slices.Delete(fv.logic.Temp_section.Atachements, i, i+1)
				fv.atachement_list.Refresh()
			}
		},
	)
	fv.form = widget.NewForm(
		fv.items["title"],
		fv.items["start"],
		fv.items["end"],
		fv.items["description"],
		fv.items["color"],
		fv.items["Atachement"],
	)
}

func (fv *FormView) GetCanvas() fyne.CanvasObject {
	add := widget.NewButton("Add Section", func() {
		section, err := logic.NewSection(
			fv.items["title"].Widget.(*widget.Entry).Text,
			fv.items["start"].Widget.(*timeEntry).Text,
			fv.items["end"].Widget.(*timeEntry).Text,
			fv.items["description"].Widget.(*widget.Entry).Text,
			fv.logic.Temp_section.Color,
			fv.logic.Temp_section.Atachements,
		)
		if err != nil {
			fmt.Printf("error while creating section %v", err)
		}
		fv.logic.Temp_section = section
		fv.logic.AddSectionToSplit()
		fv.logic.Temp_section = logic.NewStubSection()
		fv.items["title"].Widget.(*widget.Entry).Text = ""
		fv.items["start"].Widget.(*timeEntry).Text = ""
		fv.items["end"].Widget.(*timeEntry).Text = ""
		fv.items["description"].Widget.(*widget.Entry).Text = ""
		fv.logic.Temp_section.Color = color.RGBA{}
		fv.logic.Temp_section.Atachements = make([]string, 0)
		fv.split_widget.Refresh()
		fv.form.Refresh()
	})
	submit := widget.NewButton("Submit", func() {
		err := fv.logic.Submit()
		if err != nil {
			dialog.ShowError(err, fv.router.Window)
			return
		}
		fv.logic.Log()
		fmt.Println("im here form view button")
		fv.router.Sidebar = NewSideBar(fv.router)
		fv.router.ShowEmpty()
	})
	buttons := container.NewVBox(add, submit)
	view := container.NewBorder(fv.form, buttons, nil, nil, fv.atachement_list)
	return view
}
