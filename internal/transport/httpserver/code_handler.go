package httpserver

import (
	"net/http"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/antongoncharik/sso/internal/service"
	"github.com/gin-gonic/gin"
)

type CodeHandler struct {
	svc *service.Service
}

func NewCodeHandler(svc *service.Service) *CodeHandler {
	return &CodeHandler{svc}
}

func (h *CodeHandler) ExchangeCode(ctx *gin.Context) {
	var exchangeCode domain.ExchangeCode

	err := ctx.ShouldBindJSON(&exchangeCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.svc.Code.ExchangeCode(exchangeCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
