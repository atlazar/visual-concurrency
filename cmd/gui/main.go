package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/atlazar/visual-concurrency/internal/gui"
)

func main() {
	a := app.New()

	w := a.NewWindow("VisualConcurrency")
	view := gui.NewMainView()
	w.SetContent(view.CompList)
	w.ShowAndRun()
}
