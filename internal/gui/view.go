package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainView struct {
	Container *fyne.Container
}

func NewMainView() *MainView {
	return &MainView{
		Container: container.New(
			layout.NewVBoxLayout(),
			widget.NewLabel("counter 1"),
			widget.NewLabel("counter 2"),
			widget.NewButton("Open", func() {}),
		),
	}
}
