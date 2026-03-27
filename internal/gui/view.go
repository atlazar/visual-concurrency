package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type MainView struct {
	CompList *widget.List
}

func NewMainView() *MainView {
	compList := widget.NewList(
		func() int {
			return 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("not started")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText("hello")
		},
	)
	return &MainView{CompList: compList}
}
