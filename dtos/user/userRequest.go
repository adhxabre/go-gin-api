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
	OldPassword string `json:"old_password" form:"old_password" validate:"required"`
	NewPassword string `json:"new_password" form:"new_password" validate:"required"`
}
