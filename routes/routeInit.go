package routes

import (
	routesdebug "go-gin-api/routes/debug"
	routesv1 "go-gin-api/routes/v1"

	"github.com/gin-gonic/gin"
)

func RouteInit(e *gin.Engine) {
	routesv1.V1(e)
	routesdebug.Debug(e)
}
