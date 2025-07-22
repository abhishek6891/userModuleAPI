package login

import (
	"awesomeProject/src/user/userLocalDb"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone"`
	Password string `json:"password" binding:"required"`
}

func LoginRequestWithPost(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var loginUser userLocalDb.LoginUser // This struct should probably be defined in userLocalDb

	// ValidUsers: it holds all allowed email addresses & passwords
	passwd, isExist := userLocalDb.ValidUsers[req.Email]
	if isExist {
		if passwd == req.Password {
			loginUser = userLocalDb.LoginUser{Email: req.Email, Password: req.Password, Phone: req.Phone, Status: true, Message: "Logged in successfully."}
			userLocalDb.LoggedInUserList = append(userLocalDb.LoggedInUserList, loginUser)
			c.JSON(http.StatusOK, gin.H{"data": loginUser})
		} else {
			loginUser = userLocalDb.LoginUser{Email: req.Email, Status: false, Message: "Password is incorrect."}
			c.JSON(http.StatusUnauthorized, gin.H{"data": loginUser}) // Use 401 for incorrect password
		}
		return
	} else {
		loginUser = userLocalDb.LoginUser{Email: req.Email, Status: false, Message: "Login is not authorized."}
		c.JSON(http.StatusUnauthorized, gin.H{"data": loginUser}) // Use 401 for unauthorized login
	}
}
