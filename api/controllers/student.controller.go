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

func NewStudentController(studentService services.StudentService) StudentController {
	return StudentController{
		StudentService: studentService,
	}
}

func (controller *StudentController) CreateStudent(ctx *gin.Context) {
	var student models.Student

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.StudentService.CreateStudent(student, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *StudentController) GetStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	student, err := controller.StudentService.GetStudent(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (controller *StudentController) GetAll(ctx *gin.Context) {
	students, err := controller.StudentService.GetAllStudents(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}

func (controller *StudentController) UpdateStudent(ctx *gin.Context) {
	var student models.UpdateStudent

	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.StudentService.UpdateStudent(&objectId, student, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *StudentController) DeleteStudent(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := controller.StudentService.DeleteStudent(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *StudentController) RegisterStudentRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/students", controller.CreateStudent)
	routerGroup.GET("/students/:id", controller.GetStudent)
	routerGroup.GET("/students", controller.GetAll)
	routerGroup.PATCH("/students/:id", controller.UpdateStudent)
	routerGroup.DELETE("/students/:id", controller.DeleteStudent)
}
