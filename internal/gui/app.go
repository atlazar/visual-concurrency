package gui

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"github.com/atlazar/visual-concurrency/internal/gui/model"
	"github.com/atlazar/visual-concurrency/internal/gui/presenter"
	"github.com/atlazar/visual-concurrency/internal/gui/view"
)

type App struct {
	view      *view.Counter
	presenter *presenter.Counter
}

func NewApp() *App {
	counterView := view.NewCounterView()
	counterModel := model.NewCounterModel()
	return &App{
		view:      counterView,
		presenter: presenter.NewCounterPresenter(counterView, counterModel),
	}
}

func (app *App) Run() {
	a := fyneApp.New()

	w := a.NewWindow("VisualConcurrency")
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(300, 10))
	app.view.ShowAt(w)
}
