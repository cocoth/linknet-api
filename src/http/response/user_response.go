package response

type UserResponse struct {
	Id    string        `json:"id"`
	Name  string        `json:"name"`
	Email string        `json:"email"`
	Phone string        `json:"phone"`
	Role  *RoleResponse `json:"role"`
}

type RoleResponse struct {
	Name string `json:"name"`
}

type RegisterUserResponse struct {
	Id string `json:"id"`
}

type LoginUserResponse struct {
	Id    string `json:"user_id"`
	Token string `json:"token"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}
