package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

// Setup Gin Router
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	return r
}

func main() {
	appPort, isSet := os.LookupEnv("APP_PORT")
	if !isSet {
		appPort = "8080"
	}

	r := setupRouter()
	r.Run(fmt.Sprintf("0.0.0.0:%s", appPort))
}
