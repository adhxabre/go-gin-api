package handlers

import (
	dtos "go-gin-api/dtos/result"
	usersdtos "go-gin-api/dtos/user"
	"go-gin-api/models"
	"go-gin-api/pkg/bcrypt"
	"go-gin-api/repositories"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type handler struct {
	UserRepositories repositories.UserRepositories
}

type updateResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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

func (h *handler) UpdateName(c *gin.Context) {
	var err error

	request := new(usersdtos.UpdateNameRequest)
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepositories.GetUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if request.Name != "" {
		user.Name = request.Name
	} else if request.Name == "" {
		c.JSON(1001, dtos.ErrorResult{Code: 1001, Message: "Your name is empty."})
		return
	}

	_, err = h.UserRepositories.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateResponse{Code: http.StatusOK, Message: "Name updated successfully"})
}

func (h *handler) UpdateEmail(c *gin.Context) {
	var err error

	request := new(usersdtos.UpdateEmailRequest)
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepositories.GetUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if request.Email != "" {
		user.Email = request.Email
	} else if request.Email == "" {
		c.JSON(1002, dtos.ErrorResult{Code: 1002, Message: "Your email is empty."})
		return
	}

	_, err = h.UserRepositories.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateResponse{Code: http.StatusOK, Message: "Email successfully updated."})
}

func (h *handler) UpdatePassword(c *gin.Context) {
	var err error

	request := new(usersdtos.UpdatePasswordRequest)
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepositories.GetUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if request.OldPassword != "" {
		isValid := bcrypt.CheckHashedPassword(request.OldPassword, user.Password)
		if !isValid {
			c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: "Your old password is wrong."})
			return
		}

		if request.NewPassword != "" {
			password, err := bcrypt.HashingPassword(request.NewPassword)
			if err != nil {
				c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
				return
			}

			user.Password = password
		} else if request.NewPassword == "" {
			c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: 1004, Message: "Your new password is empty."})
			return
		}
	} else if request.OldPassword == "" {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: 1003, Message: "Your old password is empty."})
		return
	}

	_, err = h.UserRepositories.UpdateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateResponse{Code: http.StatusOK, Message: "Password successfully updated."})
}

func (h *handler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepositories.GetUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data, err := h.UserRepositories.DeleteUser(user, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.SuccessResult{Code: http.StatusOK, Data: data})
}

func convertResponse(u models.User) usersdtos.UserResponse {
	return usersdtos.UserResponse{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}
