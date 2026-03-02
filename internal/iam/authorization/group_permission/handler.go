package group_permission

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

	groupPermissions, pagination, err := h.service.List(c.Request.Context(), request.Page)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtos.ReplyDTO{
		Code:       dtos.SUCCESS,
		Data:       groupPermissions,
		Pagination: pagination,
		Message:    "group_permission.listed",
	})
}

func (h *Handler) Details(c *gin.Context) {
	var request DetailsRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(err)
		return
	}

	groupPermission, err := h.service.Details(c.Request.Context(), request.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtos.ReplyDTO{
		Code:    dtos.SUCCESS,
		Data:    groupPermission,
		Message: "group_permission.details",
	})
}

func (h *Handler) Create(c *gin.Context) {
	var request CreateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	err := h.service.Create(c.Request.Context(), &request.GroupPermission)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, dtos.ReplyDTO{
		Code:    dtos.SUCCESS,
		Message: "group_permission.created",
	})
}

func (h *Handler) Update(c *gin.Context) {
	var request UpdateRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	err := h.service.Update(c.Request.Context(), &request.GroupPermission)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtos.ReplyDTO{
		Code:    dtos.SUCCESS,
		Message: "group_permission.updated",
	})
}

func (h *Handler) Delete(c *gin.Context) {
	var request DeleteRequest
	if err := c.ShouldBindUri(&request); err != nil {
		c.Error(err)
		return
	}

	err := h.service.Delete(c.Request.Context(), request.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, dtos.ReplyDTO{
		Code:    dtos.SUCCESS,
		Message: "group_permission.deleted",
	})
}
