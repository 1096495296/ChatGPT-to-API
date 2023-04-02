package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/r3labs/sse/v2"
)

var HOST string
var PORT string
var PUID string
var ACCESS_TOKENS []string

func init() {
	HOST = os.Getenv("SERVER_HOST")
	PORT = os.Getenv("SERVER_PORT")
	PUID = os.Getenv("PUID")
	if HOST == "" {
		HOST = "127.0.0.1"
	}
	if PORT == "" {
		PORT = "8080"
	}
}

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	/// Admin routes
	router.PATCH("/admin/puid", admin_check, func(c *gin.Context) {
		// Get the PUID from the request and update the PUID
		puid := c.Query("puid")
		if puid != "" {
			PUID = puid
		} else {
			c.String(400, "puid not provided")
			return
		}
		c.String(200, "puid updated")

	})
	router.PATCH("/admin/password", admin_check, func(c *gin.Context) {
		// Get the password from the request and update the password
		password := c.Query("password")
		if password != "" {
			ADMIN_PASSWORD = password
		} else {
			c.String(400, "password not provided")
			return
		}
		c.String(200, "password updated")
	})
	router.PATCH("/admin/tokens", admin_check, func(c *gin.Context) {
		// Get the tokens from the request (json) and update the tokens
		var tokens []string
		err := c.BindJSON(&tokens)
		if err != nil {
			c.String(400, "tokens not provided")
			return
		}
		ACCESS_TOKENS = tokens
		c.String(200, "tokens updated")
	})
	router.Run(HOST + ":" + PORT)
}
