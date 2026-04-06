package presenter

type View interface {
	SetOnButtonClick(func())
	SetOneLabelText(text string)
	SetTwoLabelText(text string)
	UpdateButton(text string, active bool)
}

type Model interface {
	SetCounterOneHandler(func(string))
	SetCounterTwoHandler(func(string))
	SetFinishHandler(func())
	StartCount()
	StopCount()
}

type Counter struct {
	view  View
	model Model
	state State
}

func NewCounterPresenter(view View, model Model) *Counter {
	presenter := Counter{
		view:  view,
		model: model,
	}
	presenter.SetState(NotStarted)
	view.SetOnButtonClick(presenter.onButtonClick)

	model.SetCounterOneHandler(presenter.onCounterOneChanged)
	model.SetCounterTwoHandler(presenter.onCounterTwoChanged)
	model.SetFinishHandler(presenter.onFinish)
	return &presenter
}

func (p *Counter) onButtonClick() {
	switch p.state {
	case NotStarted:
		p.SetState(Started)
		p.model.StartCount()
	case Started:
		p.SetState(Stopping)
		p.model.StopCount()
	}
}

func (p *Counter) onCounterOneChanged(text string) {
	p.view.SetOneLabelText(text)
}

func (p *Counter) onCounterTwoChanged(text string) {
	p.view.SetTwoLabelText(text)
}

func (p *Counter) onFinish() {
	p.SetState(NotStarted)
}

func (p *Counter) SetState(state State) {
	p.state = state
	p.view.UpdateButton(p.state.ButtonText, p.state.Active)
	labelText := p.state.LabelText
	if len(labelText) > 0 {
		p.view.SetOneLabelText(labelText)
		p.view.SetTwoLabelText(labelText)
	}
}
