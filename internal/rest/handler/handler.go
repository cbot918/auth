package handler

import (
	"auth/internal/db"
	"auth/internal/util"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func Handle(c *gin.Context) {
// 	c.String(200, "/api/auth")
// 	util.Logg("/api/auth")
// }

// func Handle2(c *gin.Context) {
// 	util.Logg("c.FullPath: " + c.FullPath())
// 	name := c.Param("name")
// 	util.Logg("c.Param: " + name)
// 	email := c.Query("email")
// 	util.Logg("c.Query: " + email)
// 	c.String(http.StatusOK, "%s %s", name, email)
// }

type Handler struct {
	db *db.DB
}

func NewHandler(db *db.DB) *Handler {
	obj := new(Handler)
	obj.db = db
	return obj
}

func (h *Handler) Ping(c *gin.Context) {
	c.String(200, "pong")
	fmt.Println("[*] http request /ping")
}

func (h *Handler) Script(c *gin.Context) {
	c.String(200, fmt.Sprintf("%s.sh", c.Param("cname")))
	fmt.Println("[*] http request /m1/script")
}

func (h *Handler) Signup(c *gin.Context) {

	// operate db
	h.db.CreateRow("yale", "yale918", "12345")
	// h.db.ReadTable()
	h.db.ReadRow(1)

	message, err := ioutil.ReadAll(c.Request.Body)
	util.Checke(err, "read request body failed")

	c.String(http.StatusOK, string(message))
	// util.Logg(string(message))

}

func (h *Handler) Login(c *gin.Context) {
	message, err := ioutil.ReadAll(c.Request.Body)
	util.Checke(err, "read request body failed")

	c.String(http.StatusOK, string(message))
	util.Logg(string(message))
}

func (h *Handler) Logout(c *gin.Context) {

}

func (h *Handler) Forgetpw(c *gin.Context) {

}
