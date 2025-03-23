package main

import (
	"log"
	"os"

	"pos-backend/config"
	"pos-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize the database
	config.InitDB()

	// Setup Gin router
	r := gin.Default()

	// Register routes
	routes.UserRoutes(r)
	routes.ProductRoutes(r)
	routes.SaleRoutes(r)

	// Register Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("POS Backend running on port", port)
	r.Run(":" + port)
}
