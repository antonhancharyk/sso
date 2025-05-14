package http

import (
	"net/http"

	"github.com/antongoncharik/sso/internal/api/http/handler"
	"github.com/antongoncharik/sso/internal/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func GetRoutes(hdl *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/healthz", healthz)

	r.Use(middleware.UseCORS())

	r.Static("/static", "./static")
	r.LoadHTMLGlob("./static/templates/*")

	r.GET("/register", hdl.ShowRegisterForm)
	r.POST("/register", hdl.RegisterForm)
	r.GET("/login", hdl.ShowLoginForm)
	r.POST("/login", hdl.LoginForm)
	r.POST("/exchange", hdl.ExchangeCode)
	r.POST("/refresh", hdl.RefreshToken)
	r.GET("/validate", hdl.ValidateToken)

	return r
}

func healthz(ctx *gin.Context) {
	ctx.String(http.StatusOK, "ok")
}
