package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"qualifighting.backend.de/api/controllers"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/lib"
)

var (
	server            *gin.Engine
	studentService    services.StudentService
	studentController controllers.StudentController
	ctx               context.Context
	studentCollection *mongo.Collection
)

func init() {
	ctx = context.TODO()
	studentCollection = lib.MongoDBStudentCollection()
	studentService = services.NewStudentService(studentCollection, ctx)
	studentController = controllers.New(studentService)
	server = gin.Default()
}

func main() {
	// TODO: move routes to routes packages
	basepath := server.Group("/v1")
	studentController.RegisterStudentRoutes(basepath)

	log.Fatal(server.Run(":8080"))

	// router := routes.NewRouter()
	// port := fmt.Sprintf(":%d", lib.GetAppConfig().Port)
	// router.Run(port)

}
