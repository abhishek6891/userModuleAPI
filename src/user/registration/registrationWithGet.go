package registration

import (
	"awesomeProject/src/user/userLocalDb"
	"awesomeProject/src/user/userRegistrationDb"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func RegistrationRequestWithGet(c *gin.Context) {
	query := c.Request.URL.RawQuery
	params, _ := url.ParseQuery(query)
	username := params.Get("username")
	email := params.Get("email")
	phone := params.Get("phone")
	password := params.Get("password")
	fullname := params.Get("fullname")
	gender := params.Get("gender")
	dob := params.Get("dob")

	var registrationUser userRegistrationDb.RegistrationUser
	//_, isExist := userRegistrationDb.TrueUser[email]
	_, isExist := userLocalDb.ValidUsers[email]
	if isExist {
		registrationUser = userRegistrationDb.RegistrationUser{Username: username, Email: email, Password: password, Phone: phone, FullName: fullname, Gender: gender, DOB: dob, Status: true, Message: " Registered successfully."}
		userRegistrationDb.RegisteredInUserList = append(userRegistrationDb.RegisteredInUserList, registrationUser)
		c.JSON(http.StatusOK, gin.H{"data": registrationUser})
		return
	} else {
		var errorMsg = Error{Status: false, Message: "This is Demo practice, Only pre defined emails are allowed."}
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": errorMsg})
	}
}

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}
