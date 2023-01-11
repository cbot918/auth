package rest

import (
	"auth/internal/rest/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
}

func NewRoute() *Route {
	r := new(Route)
	return r
}

func (r *Route) RegistRoute(router *gin.Engine) *gin.Engine {
	// router.GET("/api/auth", h.Handle)
	// router.GET("/api/auth/:name", h.Handle2)

	h := handler.NewHandler()

	// ping
	router.GET("/ping", h.Ping)

	// script
	router.GET("/m1/script/:cname", h.Script)

	// auth
	router.POST("/m1/signup", h.Signup)
	router.POST("/m1/login", h.Login)
	router.GET("/m1/logout", h.Logout)
	router.GET("/m1/forgetpw", h.Forgetpw)
	return router
}
