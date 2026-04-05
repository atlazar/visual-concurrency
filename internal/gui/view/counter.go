package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Counter struct {
	oneLabel  *widget.Label
	twoLabel  *widget.Label
	button    *widget.Button
	container *fyne.Container
}

func NewCounterView() *Counter {
	oneLabel := widget.NewLabel("")
	twoLabel := widget.NewLabel("")
	button := widget.NewButton("Count", func() {})
	return &Counter{
		oneLabel: oneLabel,
		twoLabel: twoLabel,
		button:   button,
		container: container.New(
			layout.NewVBoxLayout(),
			oneLabel,
			twoLabel,
			button,
		),
	}
}

func (c *Counter) ShowAt(w fyne.Window) {
	w.SetContent(c.container)
	w.ShowAndRun()
}

func (c *Counter) SetOnButtonClick(onClick func()) {
	c.button.OnTapped = onClick
}

func (c *Counter) SetOneLabelText(text string) {
	go fyne.Do(func() {
		c.oneLabel.SetText(text)
		c.oneLabel.Refresh()
	})
}

func (c *Counter) SetTwoLabelText(text string) {
	go fyne.Do(func() {
		c.twoLabel.SetText(text)
		c.twoLabel.Refresh()
	})
}
