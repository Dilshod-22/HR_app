package delivery

import (
	"HR/internal/domain"
	"HR/internal/usecase"
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetEmployee(c *gin.Context) {
	data, err := usecase.GetEmployee()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func CreateEmployee(c *gin.Context) {
	var emp domain.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := usecase.CreateEmployee(emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Yaratildi!"})
}

func UpdateEmployee(c *gin.Context) {
	id := c.Param("id")
	var emp domain.Employee
	if err := c.ShouldBindJSON(&emp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := usecase.UpdateEmployee(id, emp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Yangilandi!"})
}

func DeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := usecase.DeleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "O'chirildi!"})
}

func SoftDeleteEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := usecase.SoftDeleteEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Vaqtincha o'chirildi!"})
}

func RecoverEmployee(c *gin.Context) {
	id := c.Param("id")
	if err := usecase.RecoverEmployee(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tiklandi!"})
}

func HashMD5(c *gin.Context) {
	var body struct {
		KeyValue string `json:"key_value" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hash := fmt.Sprintf("%x", md5.Sum([]byte(body.KeyValue)))
	c.JSON(http.StatusOK, gin.H{"md5": hash})
}
