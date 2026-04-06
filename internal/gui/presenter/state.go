package presenter

type labelState string
type buttonState string

var (
	labelNotStarted labelState = "not started"
	labelStarting   labelState = "starting"
	labelStarted    labelState = "started"
	labelCancelled  labelState = "cancelled"

	buttonNotStarted buttonState = "not started"
	buttonStarted    buttonState = "started"
	buttonStopping   buttonState = "stopping"
)
