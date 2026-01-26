package resthandler

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	log *slog.Logger
}

func NewHandler(log *slog.Logger) *RestHandler {
	return &RestHandler{
		log: log,
	}
}

func (h *RestHandler) Create(ctx *gin.Context) {
	panic("not implemented")
}

func (h *RestHandler) Delete(ctx *gin.Context) {
	panic("not implemented")
}

func (h *RestHandler) GetAll(ctx *gin.Context) {
	panic("not implemented")
}
