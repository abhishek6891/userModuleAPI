package scholar

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ScholarData struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Dob    string `json:"dob"`
	Gender string `json:"gender"`
}

func ScholarRequest(c *gin.Context) {

	scholarData := ScholarData{Name: "abhiishekkkkk", Email: "abhisjek@gmail,.com", Phone: "7376039833", Dob: "15/02/2003", Gender: "Male"}
	c.JSON(http.StatusOK, gin.H{"data": scholarData})

}
