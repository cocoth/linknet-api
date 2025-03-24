package response

type RegisterUserResponse struct {
	ID string `json:"id"`
}

type LoginUserResponse struct {
	ID           string `json:"user_id"`
	SessionToken string `json:"session_token"`
	CsrfToken    string `json:"csrf_token"`
}
