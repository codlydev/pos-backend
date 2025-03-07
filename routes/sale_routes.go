package routes

import (
	"pos-backend/controllers"
	"pos-backend/middlewares"

	"github.com/gin-gonic/gin"
)

func SaleRoutes(r *gin.Engine) {
	r.GET("/sales", middlewares.AuthMiddleware(), controllers.GetSales)
	r.POST("/sales", middlewares.AuthMiddleware(), controllers.CreateSale)
}
