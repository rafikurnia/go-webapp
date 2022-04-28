package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	r.Run("0.0.0.0:8080")
}
