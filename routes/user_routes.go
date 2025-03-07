package routes

import "github.com/gin-gonic/gin"

func UserRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
}
