package main

import (
	"usermanager/model"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Println(">>Main")
	log.Info("TESTING")
	log.WithFields(log.Fields{
		"name":    "anand",
		"number":  333,
		"address": "tampa",
	}).Info("Here are the Fields ")
	log.SetFormatter(&log.JSONFormatter{})
	standardFields := log.Fields{
		"hostname": "staging-1",
		"appname":  "foo-app",
		"session":  "1ce3f6v",
	}
	log.WithFields(standardFields).WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1}).Info("My first ssl event from Golang")
	router := gin.Default()
	router.GET("/getAllUsers", model.GetAllUsers)
	router.GET("/healthcheck", model.GetHealthCheck)
	router.Run("localhost:8080")
}
