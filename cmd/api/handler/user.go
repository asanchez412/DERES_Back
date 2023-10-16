package handler

import (
	"errors"

	"github.com/andsanchez/DERES_Back/internal/user"
	"github.com/andsanchez/DERES_Back/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type loginRequest struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type signUpRequest struct {
	Username string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserHandler struct {
	userService user.Service
}

func NewUserHandler(s user.Service) *UserHandler {
	return &UserHandler{
		userService: s,
	}
}

func (h *UserHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				web.Error(c, 422, "%v", err.Error())
				return
			}
			web.Error(c, 400, "%v", err.Error())
			return
		}

		err := h.userService.Login(c, req.Username, req.Password)
		if err != nil {
			web.Error(c, 404, "%v", err.Error())
			return
		}
		web.Success(c, 200, "logged in")
	}
}

func (h *UserHandler) SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req signUpRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			var ve validator.ValidationErrors
			if errors.As(err, &ve) {
				web.Error(c, 422, "%v", err.Error())
				return
			}
			web.Error(c, 400, "%v", err.Error())
			return
		}
		err := h.userService.SignUp(c, req.Username, req.Password)
		if err != nil {
			web.Error(c, 500, "%v", err.Error())
			return
		}
		web.Success(c, 201, "user created")
	}
}
