package ui

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/core"
	"github.com/coshi-muhammd/timesplit/internal/logic"
)

type newSplitDialog struct {
	title    *widget.FormItem
	checkbox *widget.FormItem
	dialog   *dialog.CustomDialog
}
type newAtachementDialog struct {
	name   *widget.FormItem
	uri    *widget.FormItem
	temp   core.Atachement
	dialog *dialog.CustomDialog
}

func NewSplitDialauge(r *UiRouter) *newSplitDialog {
	title := widget.NewFormItem("Title:", widget.NewEntry())
	checkbox := widget.NewFormItem("Week format:", widget.NewCheck("", func(b bool) {
		b = !b
	}),
	)
	form := widget.NewForm(title, checkbox)
	popup := dialog.NewCustom("New Split", "Cancel", form, r.Window)
	form.OnSubmit = func() {
		form_logic := logic.NewFormService(r.Logic,
			title.Widget.(*widget.Entry).Text,
			checkbox.Widget.(*widget.Check).Checked,
		)
		r.ShowForm(form_logic)
		popup.Dismiss()
	}

	return &newSplitDialog{
		title:    title,
		checkbox: checkbox,
		dialog:   popup,
	}
}
func (sd *newSplitDialog) Show() {
	sd.dialog.Show()
}

func NewAtachementDialog(r *UiRouter) *newAtachementDialog {
	ad := &newAtachementDialog{}

	ad.name = widget.NewFormItem("Name:", widget.NewEntry())
	ad.uri = widget.NewFormItem("Path/Link:", widget.NewEntry())
	form := widget.NewForm(ad.name, ad.uri)
	ad.dialog = dialog.NewCustom("New Atachement", "Cancel", form, r.Window)
	form.OnSubmit = func() {
		var fyne_uri fyne.URI
		if strings.HasPrefix(ad.uri.Widget.(*widget.Entry).Text, "http") {
			fyne_uri, _ = storage.ParseURI(ad.uri.Widget.(*widget.Entry).Text)
		} else {
			fyne_uri = storage.NewFileURI(ad.uri.Widget.(*widget.Entry).Text)
		}
		fmt.Println(fyne_uri)
		ad.temp = logic.NewAtachement(ad.name.Widget.(*widget.Entry).Text, fyne_uri.String())
		ad.dialog.Dismiss()
	}
	return ad
}
func (na *newAtachementDialog) Show() {
	na.dialog.Show()
}
