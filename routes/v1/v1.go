package routesv1

import (
	"github.com/gin-gonic/gin"
)

func V1(e *gin.Engine) {
	routes := e.Group("/v1")
	{
		UserRoutes(routes)
	}
}
