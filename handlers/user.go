package handlers

import (
	dtos "go-gin-api/dtos/result"
	usersdtos "go-gin-api/dtos/user"
	"go-gin-api/models"
	"go-gin-api/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	UserRepositories repositories.UserRepositories
}

func HandlerUser(UserRepositories repositories.UserRepositories) *handler {
	return &handler{UserRepositories}
}

func (h *handler) FindUsers(c *gin.Context) {
	users, err := h.UserRepositories.FindUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	var userResponses []usersdtos.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, convertResponse(user))
	}

	if len(userResponses) == 0 {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResult{Code: http.StatusInternalServerError, Message: "Data is missing or not yet inputted, create one!"})
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResult{Code: http.StatusOK, Data: userResponses})
}

func (h *handler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepositories.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResult{Code: http.StatusOK, Data: user})
}

func convertResponse(u models.User) usersdtos.UserResponse {
	return usersdtos.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
