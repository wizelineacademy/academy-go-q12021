package controller

// Context represents the context of the current HTTP request. It holds request and
// response objects, path, path parameters, data and registered handler.
type Context interface {
	JSON(code int, i interface{}) error
	Bind(i interface{}) error
}
