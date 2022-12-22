package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type ExamController struct {
	ExamService services.ExamService
}

func NewExamController(examService services.ExamService) ExamController {
	return ExamController{
		ExamService: examService,
	}
}

func (controller *ExamController) CreateExam(ctx *gin.Context) {
	var exam models.Exam

	if err := ctx.ShouldBindJSON(&exam); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.ExamService.CreateExam(&exam, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *ExamController) GetExam(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	exam, err := controller.ExamService.GetExam(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exam)
}

func (controller *ExamController) GetAllExams(ctx *gin.Context) {
	exams, err := controller.ExamService.GetAllExams(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, exams)
}

func (controller *ExamController) UpdateExams(ctx *gin.Context) {
	var exam models.UpdateExam

	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	if err := ctx.ShouldBindJSON(&exam); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.ExamService.UpdateExam(&objectId, &exam, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *ExamController) DeleteExam(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := controller.ExamService.DeleteExam(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *ExamController) RegisterExamRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/exams", controller.CreateExam)
	routerGroup.GET("/exams/:id", controller.GetExam)
	routerGroup.GET("/exams", controller.GetAllExams)
	routerGroup.PATCH("/exams/:id", controller.UpdateExams)
	routerGroup.DELETE("/exams/:id", controller.DeleteExam)
}
