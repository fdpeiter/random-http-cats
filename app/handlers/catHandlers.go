package handlers

import (
	"github.com/fdpeiter/random-http-cats/app/logger"
	"github.com/fdpeiter/random-http-cats/app/services"
	"github.com/gin-gonic/gin"
)

func RandomCatHandler(c *gin.Context) {
	catServiceLogger := logger.NewLogger("CatService")
	catService := services.NewCatService(catServiceLogger)
	catService.ServeRandomCat(c)
}

func SpecificCatHandler(c *gin.Context) {
	statusCode := c.Param("statusCode")
	catServiceLogger := logger.NewLogger("CatService")
	catService := services.NewCatService(catServiceLogger)
	catService.ServeSpecificCat(c, statusCode)
}
