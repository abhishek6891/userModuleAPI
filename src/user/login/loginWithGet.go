package login

import (
	"awesomeProject/src/user/userLocalDb"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func LoginRequestWithGet(c *gin.Context) {
	query := c.Request.URL.RawQuery
	params, _ := url.ParseQuery(query)
	email := params.Get("email")
	phone := params.Get("phone")
	password := params.Get("password")

	var loginUser userLocalDb.LoginUser
	_, isExist := userLocalDb.ValidUsers[email]
	if isExist {
		loginUser = userLocalDb.LoginUser{Email: email, Password: password, Phone: phone, Status: true, Message: "Logged in successfully."}
		userLocalDb.LoggedInUserList = append(userLocalDb.LoggedInUserList, loginUser)
		c.JSON(http.StatusOK, gin.H{"data": loginUser})
		return
	} else {
		loginUser = userLocalDb.LoginUser{Email: email, Status: false, Message: "Login is not authorized."}
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{"data": loginUser})
	}
}
