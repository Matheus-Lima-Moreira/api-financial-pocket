package organizations

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

	organizations, pagination, err := h.service.List(c.Request.Context(), request.Page)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtos.ReplyDTO{
		Data:       organizations,
		Pagination: pagination,
		Message:    "organization.listed",
	})
}

func (h *Handler) Details(c *gin.Context) {
	var request DetailsRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(err)
		return
	}

	organization, err := h.service.Details(c.Request.Context(), request.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtos.ReplyDTO{
		Data:    organization,
		Message: "organization.details",
	})
}

func (h *Handler) Create(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	err := h.service.Create(c.Request.Context(), &request.Organization)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dtos.ReplyDTO{
		Message: "organization.created",
	})
}

func (h *Handler) Update(c *gin.Context) {
	var request UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	err := h.service.Update(c.Request.Context(), &request.Organization)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, dtos.ReplyDTO{
		Message: "organization.updated",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	var input DeleteRequest
	if err := c.ShouldBindUri(&input); err != nil {
		c.Error(err)
		return
	}
}
