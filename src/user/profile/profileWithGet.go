package profile

import (
	"awesomeProject/src/user/profileDb"
	"awesomeProject/src/user/userLocalDb"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func ProfileRequestWithGet(c *gin.Context) {
	query := c.Request.URL.RawQuery
	params, _ := url.ParseQuery(query)
	name := params.Get("name")
	email := params.Get("email")
	phone := params.Get("phone")
	gender := params.Get("gender")
	dob := params.Get("dob")
	about := params.Get("about")
	profilePhoto := params.Get("profilePhoto")
	address := params.Get("address")
	location := params.Get("location")

	var profileUser profileDb.ProfileUser
	//_, isExist := profileDb.TrueUser[email]
	_, isExist := userLocalDb.ValidUsers[email]
	if isExist {
		profileUser = profileDb.ProfileUser{Name: name, Email: email, Phone: phone, Gender: gender, Dob: dob, About: about, ProfilePhoto: profilePhoto, Address: address, Location: location, Status: true, Message: " Profile created successfully."}
		profileDb.ProfileList = append(profileDb.ProfileList, profileUser)
		c.JSON(http.StatusOK, gin.H{"data": profileUser})
		return
	} else {
		profileUser = profileDb.ProfileUser{Email: email, Status: false, Message: "User profile does not exist."}
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": profileUser})
	}
}
