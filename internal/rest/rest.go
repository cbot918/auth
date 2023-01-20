package rest

import (
	"auth/internal/db"
	"auth/internal/rest/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	db *db.DB
}

func NewRoute(db *db.DB) *Route {
	r := new(Route)
	r.db = db
	return r
}

func (r *Route) RegistRoute(router *gin.Engine) *gin.Engine {
	// router.GET("/api/auth", h.Handle)
	// router.GET("/api/auth/:name", h.Handle2)

	h := handler.NewHandler(r.db)

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
