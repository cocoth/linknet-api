package request

type UpdateUserReq struct {
	Id       string `validate:"required, uuid" json:"id"`
	Name     string `validate:"required, min=1,max=200" json:"name"`
	Password string `validate:"min=8" json:"password"`
	Email    string `json:"email"`
}
