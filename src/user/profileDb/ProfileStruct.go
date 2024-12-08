package profileDb

type ProfileUser struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Password     string `json:"password"`
	Gender       string `json:"gender"`
	Dob          string `json:"dob"`
	About        string `json:"about"`
	ProfilePhoto string `json:"profile_photo"`
	Address      string `json:"address"`
	Location     string `json:"location"`
	Status       bool   `json:"status"`
	Message      string `json:"message"`
}
