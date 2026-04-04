package presenter

import "fmt"

type CounterView interface {
	SetOnCountClick(func())
}

type CounterModel interface{}

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
	return &presenter
}

func (p *CounterPresenter) onCountClick() {
	fmt.Println("counter clicked")
}
