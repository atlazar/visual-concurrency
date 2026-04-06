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

	labelOneState labelState
	labelTwoState labelState
	buttonState   buttonState
}

func NewCounterPresenter(view View, model Model) *Counter {
	presenter := Counter{
		view:  view,
		model: model,
	}
	presenter.setState(labelNotStarted, labelNotStarted, buttonNotStarted)
	view.SetOnButtonClick(presenter.onButtonClick)

	model.SetCounterOneHandler(presenter.onCounterOneChanged)
	model.SetCounterTwoHandler(presenter.onCounterTwoChanged)
	model.SetFinishHandler(presenter.onFinish)
	return &presenter
}

func (p *Counter) onButtonClick() {
	switch p.buttonState {
	case buttonNotStarted:
		p.setState(labelStarting, labelStarting, buttonStarted)
		p.model.StartCount()
	case buttonStarted:
		p.setState(labelCancelled, labelCancelled, buttonStopping)
		p.model.StopCount()
	}
}

func (p *Counter) onCounterOneChanged(text string) {
	p.setLabelOneState(labelStarted)
	p.view.SetOneLabelText(text)
}

func (p *Counter) onCounterTwoChanged(text string) {
	p.setLabelTwoState(labelStarted)
	p.view.SetTwoLabelText(text)
}

func (p *Counter) onFinish() {
	p.setState(labelNotStarted, labelNotStarted, buttonNotStarted)
}

func (p *Counter) setState(
	labelOneState labelState,
	labelTwoState labelState,
	buttonState buttonState,
) {
	p.setLabelOneState(labelOneState)
	p.setLabelTwoState(labelTwoState)
	p.setButtonState(buttonState)
}

func (p *Counter) setLabelOneState(state labelState) {
	if !validMove(p.labelOneState, state) {
		return
	}

	p.labelOneState = state
	labelOneText := newLabelText(state)
	if labelOneText != nil {
		p.view.SetOneLabelText(*labelOneText)
	}
}

func (p *Counter) setLabelTwoState(state labelState) {
	if !validMove(p.labelTwoState, state) {
		return
	}

	p.labelTwoState = state
	labelTwoText := newLabelText(state)
	if labelTwoText != nil {
		p.view.SetTwoLabelText(*labelTwoText)
	}
}

func (p *Counter) setButtonState(state buttonState) {
	p.buttonState = state

	switch state {
	case buttonNotStarted:
		p.view.UpdateButton("start", true)
	case buttonStarted:
		p.view.UpdateButton("stop", true)
	case buttonStopping:
		p.view.UpdateButton("stopping..", false)
	}
}

func newLabelText(state labelState) *string {
	switch state {
	case labelNotStarted:
		return new("not started")
	case labelStarting:
		return new("initializing")
	case labelCancelled:
		return new("cancelled")
	}
	return nil
}

func validMove(old labelState, new labelState) bool {
	if old == "" && new == labelNotStarted {
		return true
	}
	// to starting
	return (old == labelNotStarted && new == labelStarting) ||
		(old == labelStarted && new == labelStarting) ||
		(old == labelCancelled && new == labelStarting) ||
		// from starting
		(old == labelStarting && new == labelStarted) ||
		(old == labelStarting && new == labelCancelled)
}
