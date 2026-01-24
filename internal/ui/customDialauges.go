package ui

import (
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/coshi-muhammd/timesplit/internal/logic"
)

type newSplitDialog struct {
	title    *widget.FormItem
	checkbox *widget.FormItem
	dialog   *dialog.CustomDialog
}

func NewSplitDialauge(r *UiRouter) *newSplitDialog {
	title := widget.NewFormItem("Title:", widget.NewEntry())
	checkbox := widget.NewFormItem("Week format:", widget.NewCheck("", func(b bool) {
		b = !b
	}),
	)
	form := widget.NewForm(title, checkbox)
	popup := dialog.NewCustom("New Split", "submit", form, r.Window)
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
