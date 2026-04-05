package presenter

type State struct {
	ButtonText string
	LabelText  string
	Active     bool
}

var (
	NotStarted = State{
		ButtonText: "Start",
		Active:     true,
	}
	Started = State{
		ButtonText: "Stop",
		LabelText:  "initialize...",
		Active:     true,
	}
	Stopping = State{
		ButtonText: "Stopping",
		Active:     false,
	}
)
