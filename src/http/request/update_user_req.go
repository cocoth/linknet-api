package request

type UpdateUserRequest struct {
	Name       *string `json:"name,omitempty" validate:"omitempty,min=1,max=200"`
	Password   *string `json:"password,omitempty" validate:"min=8"`
	Email      *string `json:"email,omitempty"`
	Phone      *string `json:"phone,omitempty"`
	CallSign   *string `json:"call_sign,omitempty"`
	Contractor *string `json:"contractor,omitempty"`
	Status     *string `json:"status,omitempty"`

	Role *RoleRequest `json:"role,omitempty"`
}
