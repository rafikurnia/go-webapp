package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run("0.0.0.0:8080")
}
