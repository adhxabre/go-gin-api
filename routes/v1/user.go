package routesv1

import (
	"go-gin-api/handlers"
	"go-gin-api/pkg/database"
	"go-gin-api/repositories"

	"github.com/gin-gonic/gin"
)

func UserRoutes(e *gin.RouterGroup) {
	userRepository := repositories.RepositoryUser(database.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
}
