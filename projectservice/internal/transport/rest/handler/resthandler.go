package resthandler

import (
	"errors"
	"log/slog"
	"net/http"
	projectdomain "projectservice/internal/domain/project"
	createdto "projectservice/internal/transport/rest/handler/dto/create"
	handlmapper "projectservice/internal/transport/rest/handler/mapper"
	handlvalidator "projectservice/internal/transport/rest/handler/validator"
	createerr "projectservice/internal/usecase/error/createproject"
	"projectservice/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
)

type RestHandler struct {
	log *slog.Logger

	createProjUC interfaces.CreateProjectUsecase
}

func NewHandler(log *slog.Logger, createProjUC interfaces.CreateProjectUsecase) *RestHandler {
	return &RestHandler{
		log:          log,
		createProjUC: createProjUC,
	}
}

func (h *RestHandler) Create(ctx *gin.Context) {
	const op = "resthandler.Create"

	var userId uint32
	if val, exists := ctx.Get("userId"); exists {
		userId = val.(uint32)
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	}

	log := h.log.With(slog.String("op", op), slog.Int("userId", int(userId)))

	log.Info("start create request")

	var req *createdto.CreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Warn("error with request data", slog.String("error", err.Error()))
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

	in := handlmapper.CreateRequestToInput(req, uint32(userId))

	out, err := h.createProjUC.Execute(ctx.Request.Context(), in)
	if err != nil {
		if errors.Is(err, projectdomain.ErrInvalidName) {
			log.Info("invalid name")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if errors.Is(err, projectdomain.ErrInvalidOwnerId) {
			log.Info("invalid owner id")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else if errors.Is(err, createerr.ErrAlreadyExists) {
			log.Info("project already exists")
			ctx.JSON(http.StatusConflict, gin.H{
				"error": err.Error(),
			})
		} else {
			log.Warn("cannot create new project", slog.String("error", err.Error()))
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
		}
		return
	}

	log.Info("create request completed successfully")

	res := handlmapper.CreateOutputToResponse(out)
	ctx.JSON(http.StatusOK, res)
}

func (h *RestHandler) Delete(ctx *gin.Context) {
	panic("not implemented")
}

func (h *RestHandler) GetAll(ctx *gin.Context) {
	panic("not implemented")
}
