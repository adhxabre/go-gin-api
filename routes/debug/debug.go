package routesdebug

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Debug(e *gin.Engine) {
	e.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "Welcome to Go + Gin API app, this server build using Go and Gin, for further information, please check it on README.md"})
	})
}
