package request

type UpdateUserRequest struct {
	Id       string `validate:"required,uuid" binding:"required,uuid" json:"id"`
	Name     string `validate:"required,min=1,max=200" binding:"required,min=1,max=200" json:"name"`
	Password string `validate:"min=8" binding:"min=8" json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	RoleID   uint   `json:"role_id,omitempty"`
}
