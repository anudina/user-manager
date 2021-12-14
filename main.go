package main

import (
	"fmt"
	"strings"

	//"github.com/mattn/go-oci8"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("TEST")
	route := gin.Default()
	route.Use(setHeader())
	route.GET("getallusers", usercontroller.getAllUsers())
}

func setHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.String(), "/swaggerui/") {
			c.Header("Cache-Control", "no-cache")
		}
	}
}
