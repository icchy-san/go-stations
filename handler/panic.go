package handler

import (
	"log"
	"net/http"
)

// A PanicHandler implements for checking whether panic works or not.
type PanicHandler struct{}

// NewPanicHandler returns PanicHandler based http.Handler.
func NewPanicHandler() *PanicHandler {
	return &PanicHandler{}
}

// // ServeHTTP implements http.Handler interface.
func (h *PanicHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("panic will be occurred!")
	panic("panic occurred!")
}
