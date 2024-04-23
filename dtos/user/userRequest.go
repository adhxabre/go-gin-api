package usersdtos

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type UpdateNameRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type UpdateEmailRequest struct {
	Email string `json:"email" form:"email" validate:"required"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" form:"password" validate:"required"`
}
