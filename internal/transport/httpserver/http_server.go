package httpserver

import (
	"net/http"

	"github.com/antongoncharik/sso/internal/middleware"
	"github.com/antongoncharik/sso/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	CodeHandler  *CodeHandler
	FormHandler  *FormHandler
	TokenHandler *TokenHandler
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		CodeHandler:  NewCodeHandler(svc),
		FormHandler:  NewFormHandler(svc),
		TokenHandler: NewTokenHandler(svc),
	}
}

func GetRoutes(hdl *Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/healthz", healthz)
	r.Use(middleware.UseCORS())
	r.Static("/static", "./static")
	r.LoadHTMLGlob("./static/templates/*")
	r.GET("/register", hdl.FormHandler.ShowRegisterForm)
	r.POST("/register", hdl.FormHandler.RegisterForm)
	r.GET("/login", hdl.FormHandler.ShowLoginForm)
	r.POST("/login", hdl.FormHandler.LoginForm)
	r.POST("/exchange", hdl.CodeHandler.ExchangeCode)
	r.POST("/refresh", hdl.TokenHandler.RefreshToken)
	r.GET("/validate", hdl.TokenHandler.ValidateToken)
	return r
}

func healthz(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
