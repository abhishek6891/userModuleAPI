package userLocalDb

type LoginUser struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Status   bool   `json:"status"`
	Message  string `json:"message"`
}
