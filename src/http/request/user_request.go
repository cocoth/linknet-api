package request

type CreateUserRequest struct {
	Name     string      `validate:"required,min=1,max=200" json:"name"`
	Email    string      `json:"email"`
	Phone    string      `validate:"required" json:"phone"`
	Password string      `validate:"min=8" json:"password"`
	Role     RoleRequest `json:"role" validate:"required"`
}

type RoleRequest struct {
	Name string `validate:"required" json:"name"`
}

type RegisterUserRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	Email    string `json:"email"`
	Phone    string `validate:"required" json:"phone"`
	Password string `validate:"min=8" json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `validate:"min=8" json:"password"`
}

type LogoutUserRequest struct {
	UserID string `json:"user_id"`
}

type RefreshTokenRequest struct {
	Token string `json:"token"`
}
