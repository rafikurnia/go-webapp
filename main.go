package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rafikurnia/go-webapp/api"
	"github.com/rafikurnia/go-webapp/utils"
)

// @title         go-webapp
// @version       2.0.0
// @description   API for Contact
// @contact.name  Rafi Kurnia Putra
// @license.name  MIT License
// @license.url   https://github.com/rafikurnia/go-webapp/blob/main/LICENSE
// @BasePath      /api/v1
func main() {
	appPort, isSet := os.LookupEnv("APP_PORT")
	if !isSet {
		log.Fatal("APP_PORT is not set in the environment variable")
	}

	router, err := api.SetupRouter(gin.DebugMode, utils.DBMode)
	if err != nil {
		log.Fatalf("main() -> %v", err)
	}

	router.Run(fmt.Sprintf("0.0.0.0:%s", appPort))
}
