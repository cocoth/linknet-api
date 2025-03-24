package request

type RoleRequest struct {
	Name string `validate:"required" json:"name"`
}

type UserRequest struct {
	Name       string       `validate:"required,min=1,max=200" json:"name"`
	Email      string       `json:"email"`
	Phone      string       `validate:"required" json:"phone"`
	Password   string       `validate:"min=8" json:"password"`
	CallSign   string       `json:"call_sign"`
	Contractor *string      `json:"contractor"`
	Status     *string      `json:"status"`
	Role       *RoleRequest `json:"role" validate:"required"`
}

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

type FileUploadRequest struct {
	FileName string `json:"file_name"`
	FileUri  string `json:"file_uri"`
	FileHash string `json:"file_hash"`
	AuthorID string `json:"author_id"`
}
