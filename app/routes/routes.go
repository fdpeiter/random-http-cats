package routes

import (
	"github.com/fdpeiter/random-http-cats/app/handlers"
	"github.com/fdpeiter/random-http-cats/app/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	// Use the custom logger
	ginLogger := logger.NewLogger("gin")
	r.Use(ginLogger.GinLogger())
	r.Use(gin.Recovery())

	// Cat routes
	r.GET("/randomcat", handlers.RandomCatHandler)
	r.GET("/specificcat/:statusCode", handlers.SpecificCatHandler)

	return r
}
