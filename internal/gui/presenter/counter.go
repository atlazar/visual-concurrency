package presenter

type View interface {
	SetOnButtonClick(func())
	SetOneLabelText(text string)
	SetTwoLabelText(text string)
}

type Model interface {
	GetInitialLabel() string
	SetCounterOneHandler(func(string))
	SetCounterTwoHandler(func(string))
	Run()
}

type Counter struct {
	view  View
	model Model
}

func NewCounterPresenter(view View, model Model) *Counter {
	presenter := Counter{
		view:  view,
		model: model,
	}
	view.SetOneLabelText(model.GetInitialLabel())
	view.SetTwoLabelText(model.GetInitialLabel())
	view.SetOnButtonClick(presenter.onButtonClick)

	model.SetCounterOneHandler(presenter.onCounterOneChanged)
	model.SetCounterTwoHandler(presenter.onCounterTwoChanged)
	return &presenter
}

func (p *Counter) onButtonClick() {
	p.model.Run()
}

func (p *Counter) onCounterOneChanged(text string) {
	p.view.SetOneLabelText(text)
}

func (p *Counter) onCounterTwoChanged(text string) {
	p.view.SetTwoLabelText(text)
}
