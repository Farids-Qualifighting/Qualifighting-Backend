package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type TeacherController struct {
	TeacherService services.TeacherService
}

func NewTeacherController(teacherService services.TeacherService) TeacherController {
	return TeacherController{
		TeacherService: teacherService,
	}
}

func (controller *TeacherController) CreateTeacher(ctx *gin.Context) {
	var teacher models.Teacher

	if err := ctx.ShouldBindJSON(&teacher); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.TeacherService.CreateTeacher(&teacher, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TeacherController) GetTeacher(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	student, err := controller.TeacherService.GetTeacher(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (controller *TeacherController) GetAllTeachers(ctx *gin.Context) {
	teachers, err := controller.TeacherService.GetAllTeachers(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teachers)
}

func (controller *TeacherController) UpdateTeacher(ctx *gin.Context) {
	var teacher models.UpdateTeacher

	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	if err := ctx.ShouldBindJSON(&teacher); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.TeacherService.UpdateTeacher(&objectId, &teacher, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TeacherController) DeleteTeacher(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := controller.TeacherService.DeleteTeacher(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TeacherController) RegisterTeacherRoutes(routerGroup *gin.RouterGroup) {
	teacherRoute := routerGroup.Group("/teacher")

	teacherRoute.POST("/create", controller.CreateTeacher)
	teacherRoute.GET("/get/:id", controller.GetTeacher)
	teacherRoute.GET("/all", controller.GetAllTeachers)
	teacherRoute.PATCH("/update/:id", controller.UpdateTeacher)
	teacherRoute.DELETE("/delete/:id", controller.DeleteTeacher)
}
