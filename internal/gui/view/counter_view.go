package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type CounterView struct {
	counterOneLabel *widget.Label
	counterTwoLabel *widget.Label
	countButton     *widget.Button
	container       *fyne.Container
}

func NewCounterView() *CounterView {
	counterOneLabel := widget.NewLabel("")
	counterTwoLabel := widget.NewLabel("")
	countButton := widget.NewButton("Count", func() {})
	return &CounterView{
		counterOneLabel: counterOneLabel,
		counterTwoLabel: counterTwoLabel,
		countButton:     countButton,
		container: container.New(
			layout.NewVBoxLayout(),
			counterOneLabel,
			counterTwoLabel,
			countButton,
		),
	}
}

func (c *CounterView) ShowAt(w fyne.Window) {
	w.SetContent(c.container)
	w.ShowAndRun()
}

func (c *CounterView) SetOnCountClick(onCountClick func()) {
	c.countButton.OnTapped = onCountClick
}

func (c *CounterView) BindCounterOne(ref *string) {
	go fyne.Do(func() {
		bRef := binding.BindString(ref)
		c.counterOneLabel.Bind(bRef)
	})
}

func (c *CounterView) BindCounterTwo(ref *string) {
	go fyne.Do(func() {
		bRef := binding.BindString(ref)
		c.counterTwoLabel.Bind(bRef)
	})
}
