package action

import (
	"net/http"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/shared/dtos"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(c *gin.Context) {
	var request ListRequest
	if err := c.ShouldBindQuery(&request); err != nil {
		c.Error(err)
		return
	}

	actions, pagination, err := h.service.List(c.Request.Context(), request.Page)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtos.ReplyDTO{
		Data:       actions,
		Pagination: pagination,
		Message:    "action.listed",
	})
}
