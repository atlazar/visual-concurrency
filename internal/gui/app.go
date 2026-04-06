package gui

import (
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"github.com/atlazar/visual-concurrency/internal/gui/model"
	"github.com/atlazar/visual-concurrency/internal/gui/presenter"
	"github.com/atlazar/visual-concurrency/internal/gui/view"
)

type App struct {
	model     *model.Counter
	view      *view.Counter
	presenter *presenter.Counter

	window fyne.Window
}

func NewApp() *App {
	window := fyneApp.New().NewWindow("VisualConcurrency")
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(300, 10))

	counterView := view.NewCounterView()
	counterModel := model.NewCounterModel()
	return &App{
		model:     counterModel,
		view:      counterView,
		presenter: presenter.NewCounterPresenter(counterView, counterModel),

		window: window,
	}
}

func (app *App) Run() {
	app.view.ShowAt(app.window)
}
