package request

type FilePermRequest struct {
	UserID string `json:"user_id"`
	FileID string `json:"file_id"`
}
