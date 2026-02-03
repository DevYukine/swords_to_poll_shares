package handler

// Handler is the interface that all Discord event handlers must implement.
type Handler interface {
	GetHandlerFunc() interface{}
}
