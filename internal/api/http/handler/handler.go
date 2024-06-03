package handler

import (
	"fmt"
	"net/http"

	"github.com/antongoncharik/sso/internal/entity"
	"github.com/antongoncharik/sso/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{svc}
}

func (h *Handler) ShowRegisterForm(ctx *gin.Context) {
	var loginRegisterForm entity.LoginRegisterForm

	err := ctx.ShouldBindQuery(&loginRegisterForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "register.html", loginRegisterForm)
}

func (h *Handler) ShowLoginForm(ctx *gin.Context) {
	var loginRegisterForm entity.LoginRegisterForm

	err := ctx.ShouldBindQuery(&loginRegisterForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "login.html", loginRegisterForm)
}

func (h *Handler) RegisterForm(ctx *gin.Context) {
	var register entity.Register

	err := ctx.ShouldBind(&register)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, err := h.svc.Register(register)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redirectURL := fmt.Sprintf("%s?code=%s", register.RedirectURI, code.Code)

	ctx.Redirect(http.StatusFound, redirectURL)
}

func (h *Handler) LoginForm(ctx *gin.Context) {
	var login entity.Login

	err := ctx.ShouldBind(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, err := h.svc.Login(login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redirectURL := fmt.Sprintf("%s?code=%s", login.RedirectURI, code.Code)

	ctx.Redirect(http.StatusFound, redirectURL)
}

func (h *Handler) ExchangeCode(ctx *gin.Context) {
	var exchangeCode entity.ExchangeCode

	err := ctx.ShouldBindJSON(&exchangeCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.svc.ExchangeCode(exchangeCode)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}

func (h *Handler) ValidateToken(ctx *gin.Context) {
	var paramsToken entity.ValidateToken

	err := ctx.ShouldBindQuery(&paramsToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.svc.ValidateToken(paramsToken.Token)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (h *Handler) RefreshToken(ctx *gin.Context) {
	var paramsToken entity.ValidateToken

	token, err := h.svc.RefreshToken(paramsToken)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, token)
}
