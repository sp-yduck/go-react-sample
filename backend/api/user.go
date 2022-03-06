package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"hello": "world"})
}
