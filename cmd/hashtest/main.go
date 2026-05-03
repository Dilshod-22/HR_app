package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/hash/md5", func(c *gin.Context) {
		var body struct {
			KeyValue string `json:"key_value" binding:"required"`
		}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum([]byte(body.KeyValue)))
		c.JSON(http.StatusOK, gin.H{"md5": hash})
	})
	r.Run(":8083")
}
