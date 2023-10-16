package handler

import (
	"errors"

	"github.com/andsanchez/DERES_Back/internal/provider"
	"github.com/andsanchez/DERES_Back/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type getProviderRequest struct {
	Name string `json:"name" binding:"required"`
}

type createProviderRequest struct {
	Name string `json:"name" binding:"required"`
}

type ProviderHandler struct {
	providerService provider.Service
}

func NewProviderHandler(s provider.Service) *ProviderHandler {
	return &ProviderHandler{
		providerService: s,
	}
}

func (h *ProviderHandler) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req getProviderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				web.Error(c, 422, "%v", err.Error())
				return
			}
			web.Error(c, 400, "%v", err.Error())
			return
		}

		provider, err := h.providerService.Get(c, req.Name)
		if err != nil {
			web.Error(c, 404, "%v", err.Error())
			return
		}
		web.Success(c, 200, provider)
	}
}

func (h *ProviderHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createProviderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				web.Error(c, 422, "%v", err.Error())
				return
			}
			web.Error(c, 400, "%v", err.Error())
			return
		}
		err := h.providerService.Create(c, req.Name)
		if err != nil {
			web.Error(c, 500, "%v", err.Error())
			return
		}
		web.Success(c, 201, "provider created")
	}
}
