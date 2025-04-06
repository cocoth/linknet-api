package response

import "time"

type FilePermResponse struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	FileID    string     `json:"file_id"`
	Approved  bool       `json:"approved"`
	CreatedAt time.Time  `json:"CreatedAt"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
	DeletedAt *time.Time `json:"DeletedAt"`
}
