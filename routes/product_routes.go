package routes

import (
	"pos-backend/controllers"
	"pos-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", middlewares.AuthMiddleware(), controllers.CreateProduct)
}
