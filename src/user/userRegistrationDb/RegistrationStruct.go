package userRegistrationDb

type RegistrationUser struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Gender   string `json:"gender,omitempty"`
	DOB      string `json:"dob,omitempty"`
	Status   bool   `json:"status,omitempty"`
	Message  string `json:"message,omitempty"`
}
