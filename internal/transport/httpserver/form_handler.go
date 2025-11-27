package httpserver

import (
	"fmt"
	"net/http"

	"github.com/antongoncharik/sso/internal/domain"
	"github.com/antongoncharik/sso/internal/service"
	"github.com/gin-gonic/gin"
)

type FormHandler struct {
	svc *service.Service
}

func NewFormHandler(svc *service.Service) *FormHandler {
	return &FormHandler{svc}
}

func (h *FormHandler) ShowRegisterForm(ctx *gin.Context) {
	var loginRegisterForm domain.LoginRegisterForm

	err := ctx.ShouldBindQuery(&loginRegisterForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "register.html", loginRegisterForm)
}

func (h *FormHandler) ShowLoginForm(ctx *gin.Context) {
	var loginRegisterForm domain.LoginRegisterForm

	err := ctx.ShouldBindQuery(&loginRegisterForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "login.html", loginRegisterForm)
}

func (h *FormHandler) RegisterForm(ctx *gin.Context) {
	var register domain.Register

	err := ctx.ShouldBind(&register)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, err := h.svc.User.Register(register)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redirectURL := fmt.Sprintf("%s?code=%s", register.RedirectUri, code.Code)

	ctx.Redirect(http.StatusFound, redirectURL)
}

func (h *FormHandler) LoginForm(ctx *gin.Context) {
	var login domain.Login

	err := ctx.ShouldBind(&login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	code, err := h.svc.User.Login(login)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	redirectURL := fmt.Sprintf("%s?code=%s", login.RedirectUri, code.Code)

	ctx.Redirect(http.StatusFound, redirectURL)
}
