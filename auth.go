package main

import (
	"fmt"
	"net/http"

	"auth/internal/db"
	"auth/internal/rest"

	"auth/internal/rest/middle"

	"github.com/gin-gonic/gin"
	// "github.com/julienschmidt/httprouter"
)

func routerr(w http.ResponseWriter, r *http.Request) {
	// fmt.Printf("r.HOST: %s\n", r.Host)
	// fmt.Printf("r.METHOD: %s\n", r.Method)
	// fmt.Printf("r.RemoteAddr: %s\n", r.RemoteAddr)
	// fmt.Printf("r.RequestURI: %s\n", r.RequestURI)
	fmt.Printf("r.URL: %s\n", r.URL)
}

const (
	DB_DRIVER   = "mysql"
	DB_USER     = "root"
	DB_PASSWORD = "12345"
	DB_HOST     = "127.0.0.1"
	DB_PORT     = 8182
	DB_USE      = "cbot"
	TABLE       = "user"
)

func main() {
	router := gin.Default()
	router.Use(middle.AllowCors("*"))
	httpServer := rest.NewRoute().RegistRoute(router)

	dbconn := db.NewDB(DB_DRIVER, DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_USE)
	dbconn.SetTable(TABLE)

	/*** database CRUD ***/
	// dbconn.CreateRow("yale", "yale918")
	dbconn.ReadTable()
	// dbconn.ReadRow(1)
	// dbconn.UpdateRow("name", "node_edit", 1)
	// dbconn.DeleteRow(1)

	httpServer.Run(":8181")

}
