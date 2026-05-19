package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/VitalyCoder/todo_go/internal/core/logger"

	"go.uber.org/zap"
)

type HttpResponseHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHttpResponseHandler(
	log *core_logger.Logger,
	rw http.ResponseWriter,
) *HttpResponseHandler {
	return &HttpResponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HttpResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("write HTTP esponse", zap.Error(err))
	}
}
