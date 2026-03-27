package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/atlazar/visual-concurrency/internal/gui"
)

func main() {
	a := app.New()

	w := a.NewWindow("VisualConcurrency")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(300, 10))

	view := gui.NewMainView()
	w.SetContent(view.Container)
	w.ShowAndRun()
}
