package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type StudentController struct {
	StudentService services.StudentService
}

func New(studentService services.StudentService) StudentController {
	return StudentController{
		StudentService: studentService,
	}
}

func (controller *StudentController) CreateStudent(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.StudentService.CreateStudent(&student)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *StudentController) GetStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	student, err := controller.StudentService.GetStudent(&objectId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (controller *StudentController) GetAll(ctx *gin.Context) {
	users, err := controller.StudentService.GetAll()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (controller *StudentController) UpdateStudent(ctx *gin.Context) {
	var student models.Student

	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.StudentService.UpdateStudent(&objectId, &student)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *StudentController) DeleteStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := controller.StudentService.DeleteStudent(&objectId)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *StudentController) RegisterStudentRoutes(routerGroup *gin.RouterGroup) {
	studentRoute := routerGroup.Group("/student")

	studentRoute.POST("/create", controller.CreateStudent)
	studentRoute.GET("/get/:id", controller.GetStudent)
	studentRoute.GET("/all", controller.GetAll)
	studentRoute.PATCH("/update/:id", controller.UpdateStudent)
	studentRoute.DELETE("/delete/:id", controller.DeleteStudent)
}
