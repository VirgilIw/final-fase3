package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/virgilIw/final-fase3/internal/dto"
	pkg "github.com/virgilIw/final-fase3/pkg/jwt"
)

func AdminOnly(c *gin.Context) {
	token, isExist := c.Get("token")
	if !isExist {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Msg:     "Forbidden Access",
			Success: false,
			Data:    []any{},
			Error:   "Access Denied",
		})
		return
	}
	accessToken, ok := token.(pkg.JwtClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Data:    []any{},
			Error:   "internal server error",
		})
		return
	}
	if accessToken.Role != "admin" {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Msg:     "Forbidden Access",
			Success: false,
			Data:    []any{},
			Error:   "Access Denied",
		})
		return
	}
	c.Next()
}

func UserOnly(c *gin.Context) {
	token, isExist := c.Get("token")
	if !isExist {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Msg:     "Forbidden Access",
			Success: false,
			Data:    []any{},
			Error:   "Access Denied",
		})
		return
	}
	accessToken, ok := token.(pkg.JwtClaims)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Data:    []any{},
			Error:   "internal server error",
		})
		return
	}
	if accessToken.Role != "user" {
		c.AbortWithStatusJSON(http.StatusForbidden, dto.Response{
			Msg:     "Forbidden Access",
			Success: false,
			Data:    []any{},
			Error:   "Access Denied",
		})
		return
	}
	c.Next()
}
