package handler

import "go-api/pkg/infrastructure"

// ApplicationHTTPHandler base handler struct.
type ApplicationHTTPHandler struct {
	BaseHTTPHandler
}

// NewApplicationHTTPHandler returns ApplicationHTTPHandler instance.
func NewApplicationHTTPHandler(logger infrastructure.Logger) *ApplicationHTTPHandler {
	return &ApplicationHTTPHandler{BaseHTTPHandler: BaseHTTPHandler{Logger: logger}}
}
