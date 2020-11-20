package presenter

type ErrorPresenter interface {
	ResponseError(err error, message string)
}
