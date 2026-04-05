package presenter

type CounterView interface {
	SetOnCountClick(func())
	SetCounterOneText(text string)
	SetCounterTwoText(text string)
}

type CounterModel interface {
	GetInitialLabel() string
	SetCounterOneHandler(func(string))
	SetCounterTwoHandler(func(string))
	Run()
}

type CounterPresenter struct {
	view  CounterView
	model CounterModel
}

func NewCounterPresenter(view CounterView, model CounterModel) *CounterPresenter {
	presenter := CounterPresenter{
		view:  view,
		model: model,
	}
	view.SetCounterOneText(model.GetInitialLabel())
	view.SetCounterTwoText(model.GetInitialLabel())
	view.SetOnCountClick(presenter.onCountClick)

	model.SetCounterOneHandler(presenter.onCounterOneChanged)
	model.SetCounterTwoHandler(presenter.onCounterTwoChanged)
	return &presenter
}

func (p *CounterPresenter) onCountClick() {
	p.model.Run()
}

func (p *CounterPresenter) onCounterOneChanged(text string) {
	p.view.SetCounterOneText(text)
}

func (p *CounterPresenter) onCounterTwoChanged(text string) {
	p.view.SetCounterTwoText(text)
}
