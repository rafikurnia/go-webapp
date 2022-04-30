package api

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/rafikurnia/go-webapp/docs"
	"github.com/rafikurnia/go-webapp/models"
	"github.com/rafikurnia/go-webapp/utils"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// To throws information of unsupported method
func notAllowed(c *gin.Context) {
	utils.Throws(c, http.StatusMethodNotAllowed, errors.New("method not allowed"))
}

// Configures router of gin web framework
// It accepts mode for gin, and db that will be used to configure DB type,
// either mock or not
func SetupRouter(ginMode, dbMode string) (*gin.Engine, error) {
	if err := models.InitDB(dbMode); err != nil {
		return nil, fmt.Errorf("setupRouter(ginMode, dbMode string) -> %w", err)
	}

	gin.SetMode(ginMode)
	router := gin.Default()

	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, fmt.Errorf("setupRouter(ginMode, dbMode string) -> %w", err)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"

	contacts := &contacts{}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/contacts", contacts.getAll)
		v1.POST("/contacts", contacts.create)
		v1.PATCH("/contacts", notAllowed)
		v1.PUT("/contacts", notAllowed)
		v1.DELETE("/contacts", notAllowed)

		v1.GET("/contacts/:id", contacts.getById)
		v1.POST("/contacts/:id", notAllowed)
		v1.PATCH("/contacts:id", notAllowed)
		v1.PUT("/contacts/:id", contacts.updateById)
		v1.DELETE("/contacts/:id", contacts.deleteById)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World")
	})

	router.NoRoute(func(c *gin.Context) {
		utils.Throws(c, http.StatusNotFound, errors.New("not found"))
	})

	router.NoMethod(notAllowed)

	return router, nil
}

// Returns router that is used only for testing http calls (tests on api package)
func getRouterForTest() *gin.Engine {
	isTestOnDocker := os.Getenv("TEST_ON_DOCKER")

	var mode string
	if isTestOnDocker == "" {
		mode = utils.DBModeMock
	} else {
		mode = utils.DBModeTest
	}

	r, _ := SetupRouter(gin.TestMode, mode)
	return r
}
