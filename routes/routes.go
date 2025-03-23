package routes

import (
	"pos-backend/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func SetupRouter() *gin.Engine {
	_ = godotenv.Load() // Load environment variables
	config.InitDB()     // Initialize DB connection

	r := gin.Default()

	// Register your route groups
	UserRoutes(r)
	ProductRoutes(r)
	SaleRoutes(r)

	// Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return r
}
