package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Erros(c *gin.Context, erro error) {
	c.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      erro.Error(),
	})
}
