package response

import "time"

type RoleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserResponse struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Phone      string        `json:"phone"`
	CallSign   string        `json:"call_sign"`
	Contractor *string       `json:"contractor"`
	Status     *string       `json:"status"`
	Role       *RoleResponse `json:"role"`
	CreatedAt  time.Time     `json:"CreatedAt"`
	UpdatedAt  time.Time     `json:"UpdatedAt"`
	DeletedAt  *time.Time    `json:"DeletedAt"`
}

type FileUploadResponse struct {
	ID        string     `json:"id"`
	FileName  string     `json:"file_name"`
	FileUri   string     `json:"file_uri"`
	FileHash  string     `json:"file_hash"`
	AuthorID  *string    `json:"author_id"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}
