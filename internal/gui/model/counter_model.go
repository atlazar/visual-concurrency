package model

type CounterModel struct {
	counterOne string
	counterTwo string
}

func NewCounterModel() *CounterModel {
	return &CounterModel{
		counterOne: "one",
		counterTwo: "two",
	}
}

func (m *CounterModel) CounterOneRef() *string {
	return &m.counterOne
}

func (m *CounterModel) CounterTwoRef() *string {
	return &m.counterTwo
}
