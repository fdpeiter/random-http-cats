package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Readyz(c *gin.Context) {
	c.String(http.StatusOK, "Ready to serve")
	return
}