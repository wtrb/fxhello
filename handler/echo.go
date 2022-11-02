package handler

import (
	"io"
	"net/http"

	"go.uber.org/zap"
)

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

type EchoHandler struct {
	log *zap.Logger
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request:", zap.Error(err))
	}
}
