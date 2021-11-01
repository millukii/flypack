package logger

type ApplicationLogger interface {
	InitFunction()
	EndFunction()
	Info(msg string)
	Error(err error, msg string)
}

type ProductionLevel struct {
}

func NewApplicationLogger() (ApplicationLogger, error) {

	return &ProductionLevel{}, nil
}

func (log *ProductionLevel) Info(msg string) {

}
func (log *ProductionLevel) Error(err error, msg string) {

}
func (log *ProductionLevel) InitFunction() {

}
func (log *ProductionLevel) EndFunction() {

}