package routes

import (
	"context"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/api/controllers"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/lib"
)

var (
	studentService    services.StudentService
	studentController controllers.StudentController
	ctx               context.Context
	studentCollection *mongo.Collection
)

func NewRouter() *gin.Engine {

	ctx = context.TODO()
	studentCollection = lib.MongoDBStudentCollection()
	studentService = services.NewStudentService(studentCollection, ctx)
	studentController = controllers.New(studentService)

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "DELETE", "PUT", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	router.Use(cors.New(config))

	grp := router.Group("api/v1")
	studentController.RegisterStudentRoutes(grp)
	grp.GET("/health", controllers.Health)

	return router
}
