package ui

import "fyne.io/fyne/v2"

type View interface {
	GetCanvas() fyne.CanvasObject
}

type Component interface {
	GetCanvas() fyne.CanvasObject
}
