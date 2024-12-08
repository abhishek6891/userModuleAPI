package profileDb

type ProfileUser struct {
	Name         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	Phone        string `json:"phone,omitempty"`
	Gender       string `json:"gender,omitempty"`
	Dob          string `json:"dob,omitempty"`
	About        string `json:"about,omitempty"`
	ProfilePhoto string `json:"profile_photo,omitempty"`
	Address      string `json:"address,omitempty"`
	Location     string `json:"location,omitempty"`
	Status       bool   `json:"status,omitempty"`
	Message      string `json:"message,omitempty"`
}
