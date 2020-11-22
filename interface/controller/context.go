package controller

// Context interface serves a carrier for incoming data in Bind and Json for the response
type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
}
