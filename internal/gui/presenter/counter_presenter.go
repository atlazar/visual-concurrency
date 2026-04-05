package presenter

type CounterView interface {
	SetOnCountClick(func())
	BindCounterOne(*string)
	BindCounterTwo(*string)
}

type CounterModel interface {
	CounterOneRef() *string
	CounterTwoRef() *string
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
	view.SetOnCountClick(presenter.onCountClick)
	view.BindCounterOne(model.CounterOneRef())
	view.BindCounterTwo(model.CounterTwoRef())
	return &presenter
}

func (p *CounterPresenter) onCountClick() {
	p.model.Run()
}
