package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"qualifighting.backend.de/api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	grp := router.Group("/api/v1")

	// HEALTHCHECK
	grp.GET("/health", controllers.Health)

	return router
}
