package api

import (
	"client_service/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	client := r.Group("/client")
	{
		client.POST("/createclient", controller.CreateClientHandler)
	}
	return r
}
