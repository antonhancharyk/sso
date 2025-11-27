package httpserver

import (
	"net/http"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/antongoncharik/sso/internal/service"
	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	svc *service.Service
}

func NewTokenHandler(svc *service.Service) *TokenHandler {
	return &TokenHandler{svc}
}

func (h *TokenHandler) ValidateToken(ctx *gin.Context) {
	var paramsToken domain.ValidateToken

	err := ctx.ShouldBindQuery(&paramsToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.svc.Token.ValidateToken(paramsToken.Token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *TokenHandler) RefreshToken(ctx *gin.Context) {
	var paramsToken domain.ValidateToken

	err := ctx.ShouldBindJSON(&paramsToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.svc.Token.RefreshToken(paramsToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
