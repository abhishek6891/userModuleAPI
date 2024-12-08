package main

import (
	"awesomeProject/src/user/login"
	"awesomeProject/src/user/profile"
	"awesomeProject/src/user/registration"
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.GET("/login", login.LoginRequestWithGet)
	route.GET("/registration", registration.RegistrationRequestWithGet)
	route.GET("/profile", profile.ProfileRequestWithGet)

	route.Run() // listen and serve on 0.0.0.0:8080
}
