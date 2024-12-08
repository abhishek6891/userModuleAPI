package userRegistrationDb

type RegistrationUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Gender   string `json:"gender"`
	DOB      string `json:"dob"`
	Status   bool   `json:"status"`
	Message  string `json:"message"`
}
