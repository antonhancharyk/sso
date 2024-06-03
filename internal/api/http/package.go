package http

import (
	"fmt"

	"github.com/antongoncharik/sso/internal/api/http/handler"
	"github.com/antongoncharik/sso/internal/api/http/middleware"
	"github.com/gin-gonic/gin"
)

func RunHTTP(hdl *handler.Handler) {
	r := gin.Default()

	r.Use(middleware.UseCORS())

	r.LoadHTMLGlob("templates/*")

	r.GET("/register", hdl.ShowRegisterForm)
	r.POST("/register", hdl.RegisterForm)
	r.GET("/login", hdl.ShowLoginForm)
	r.POST("/login", hdl.LoginForm)
	r.POST("/exchange", hdl.ExchangeCode)
	r.POST("/refresh", hdl.RefreshToken)
	r.GET("/validate", hdl.ValidateToken)

	host := fmt.Sprintf(":%d", 8080)

	r.Run(host)
}
