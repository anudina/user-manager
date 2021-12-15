package model

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	Response []user `json:"data"`
}
type BaseResponse struct {
	Status   ResponseStatus `json:"status"`
	Response Response       `json:"response"`
}
type ResponseStatus struct {
	Code string `json:"code"`
	Desc string `json:"desc"`
}

type user struct {
	Name      string  `json:"name"`
	FirstName string  `json:"fname"`
	LastName  string  `json:"lname"`
	ID        float64 `json:"id"`
	Address   string  `json:"address"`
}

var getUserList = []user{
	{ID: 123, Name: "Anand", FirstName: "Anand", LastName: "Badiger", Address: "Tampa Florida"},
	{ID: 567, Name: "Rohan", FirstName: "Rohan", LastName: "B", Address: "Tampa Florida"},
	{ID: 1891, Name: "JD", FirstName: "Jayadev", LastName: "B", Address: "Tampa Florida"},
}
var getStatus = ResponseStatus{Code: "0", Desc: "Success"}

func GetAllUsers(c *gin.Context) {
	log.Info("GetAllUsers Called")
	response := BaseResponse{Status: getStatus, Response: Response{getUserList}}
	c.IndentedJSON(http.StatusOK, response)
}

func GetHealthCheck(c *gin.Context) {
	log.Info("healthCheck Called")
	c.IndentedJSON(http.StatusOK, "alive")
}
