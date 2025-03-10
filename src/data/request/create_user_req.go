package request

type CreateUserReq struct {
	Name     string  `validate:"required,min=1,max=200" json:"name"`
	Password string  `validate:"min=8" json:"password"`
	Email    string  `json:"email"`
	Role     RoleReq `json:"role" validate:"required"`
}

type RoleReq struct {
	Name string `validate:"required" json:"name"`
}
