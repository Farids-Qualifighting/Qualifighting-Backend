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
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.TeacherService.CreateTeacher(&teacher, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TeacherController) GetTeacher(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	student, err := controller.TeacherService.GetTeacher(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (controller *TeacherController) GetAllTeachers(ctx *gin.Context) {
	teachers, err := controller.TeacherService.GetAllTeachers(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, teachers)
}

func (controller *TeacherController) UpdateTeacher(ctx *gin.Context) {
	var teacher models.UpdateTeacher

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&teacher); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.TeacherService.UpdateTeacher(&objectId, &teacher, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TeacherController) DeleteTeacher(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.TeacherService.DeleteTeacher(&objectId, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TeacherController) RegisterTeacherRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/teachers", controller.CreateTeacher)
	routerGroup.GET("/teachers/:id", controller.GetTeacher)
	routerGroup.GET("/teachers", controller.GetAllTeachers)
	routerGroup.PATCH("/teachers/:id", controller.UpdateTeacher)
	routerGroup.DELETE("/teachers/:id", controller.DeleteTeacher)
}
