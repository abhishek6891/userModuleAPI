package userLocalDb

type LoginUser struct {
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Status   bool   `json:"status,omitempty"`
	Message  string `json:"message,omitempty"`
}
