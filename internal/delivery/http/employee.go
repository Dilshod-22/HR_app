package http

import (
	"HR/internal/delivery"

	"github.com/gin-gonic/gin"
)

func EmployeeRoute() *gin.Engine {
	r := gin.Default()

	r.GET("/employee", delivery.GetEmployee)
	r.POST("/employee", delivery.CreateEmployee)
	r.PUT("/employee/:id", delivery.UpdateEmployee)
	r.DELETE("/employee/:id", delivery.DeleteEmployee)
	r.POST("/recovery/:id", delivery.RecoverEmployee)
	r.DELETE("/employee/softDelete/:id", delivery.SoftDeleteEmployee)
	r.POST("/hash/md5", delivery.HashMD5)

	return r
}
