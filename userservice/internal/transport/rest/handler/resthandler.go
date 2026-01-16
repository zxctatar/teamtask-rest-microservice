package resthandler

import (
	"log/slog"
	"net/http"
	regdto "userservice/internal/transport/rest/handler/dto/registration"
	handlvalidator "userservice/internal/transport/rest/handler/validator"
	"userservice/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	log *slog.Logger

	regUC interfaces.RegistrationUsecase
}

func NewRestHandler(log *slog.Logger, regUC interfaces.RegistrationUsecase) *RestHandler {
	return &RestHandler{
		log:   log,
		regUC: regUC,
	}
}

func (h *RestHandler) Registration(ctx *gin.Context) {
	const op = "resthandler.Registration"

	var regRequest regdto.RegistrationRequest

	if err := ctx.ShouldBindJSON(&regRequest); err != nil {
		if errMap, ok := handlvalidator.MapValidationErrors(err); ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": errMap,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "bad request body",
			})
		}
		return
	}
}

func (h *RestHandler) Login(ctx *gin.Context) {
	panic("not implemented")
}
