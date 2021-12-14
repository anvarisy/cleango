package infrastuctures

import (
	"cleango/app/interfaces"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Route(connMySQL *sql.DB) *gin.Engine {
	usercontroller := interfaces.NewUserController(connMySQL)
	r := gin.Default()
	r.POST("/user", usercontroller.Save)
	r.GET("/user", usercontroller.List)
	return r
}
