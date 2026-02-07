package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgilIw/final-fase3/internal/dto"
	"github.com/virgilIw/final-fase3/internal/service"
)

// Controller itu penghubung antara HTTP request dan business logic.
// Artinya controller menangani hal-hal yang berhubungan dengan HTTP, bukan logic bisnis.

// struct

// newfuncController

// Login godoc
//
//	@Summary		User login
//	@Description	Authenticate user with email and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Login credentials"
//	@Success		200		{object}	dto.LoginResponse
//	@Failure		400		{object}	dto.ResponseError
//	@Failure		401		{object}	dto.ResponseError
//	@Router			/auth/ [post]

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register godoc
//
//	@Summary		Register new user
//	@Description	Create new user account
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.RegisterRequest	true	"Register request"
//	@Success		200		{object}	dto.Response
//	@Failure		400		{object}	dto.Response
//	@Failure		500		{object}	dto.Response
//	@Router			/auth/register [post]
func (r *AuthController) Register(c *gin.Context) {
	var req dto.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Invalid Request",
			Success: false,
			Data:    []any{},
			Error:   err.Error(),
		})
		return
	}
	err := r.authService.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Registered Failed",
			Success: false,
			Data:    []any{},
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Registered Success",
		Success: true,
		Data:    []any{},
	})
}
