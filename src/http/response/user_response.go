package response

import "time"

type UserResponse struct {
	Id         string        `json:"id"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Phone      string        `json:"phone"`
	CallSign   *string       `json:"call_sign"`
	Contractor *string       `json:"contractor"`
	Status     *string       `json:"status"`
	Role       *RoleResponse `json:"role"`
	CreatedAt  time.Time     `json:"CreatedAt"`
	UpdatedAt  time.Time     `json:"UpdatedAt"`
	DeletedAt  *time.Time    `json:"DeletedAt"`
}

type RoleResponse struct {
	Name string `json:"name"`
}

type RegisterUserResponse struct {
	Id string `json:"id"`
}

type LoginUserResponse struct {
	Id           string `json:"user_id"`
	SessionToken string `json:"session_token"`
	CsrfToken    string `json:"csrf_token"`
}

type RefreshTokenResponse struct {
	Token string `json:"token"`
}
