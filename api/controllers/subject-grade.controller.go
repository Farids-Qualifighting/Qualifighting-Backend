package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type SubjectGradeController struct {
	SubjectGradeService services.SubjectGradeService
}

func NewSubjectGradeController(subjectGradeService services.SubjectGradeService) SubjectGradeController {
	return SubjectGradeController{
		SubjectGradeService: subjectGradeService,
	}
}

func (controller *SubjectGradeController) CreateSubjectGrade(ctx *gin.Context) {
	var subjectGrade models.SubjectGrade

	if err := ctx.ShouldBindJSON(&subjectGrade); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.SubjectGradeService.CreateSubjectGrade(&subjectGrade, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SubjectGradeController) GetSubjectGrade(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	subjectGrade, err := controller.SubjectGradeService.GetSubjectGrade(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subjectGrade)
}

func (controller *SubjectGradeController) GetAllSubjectGrades(ctx *gin.Context) {
	subjectGrades, err := controller.SubjectGradeService.GetAllSubjectGrades(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subjectGrades)
}

func (controller *SubjectGradeController) UpdateSubjectGrade(ctx *gin.Context) {
	var subjectGrade models.UpdateSubjectGrade

	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	if err := ctx.ShouldBindJSON(&subjectGrade); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.SubjectGradeService.UpdateSubjectGrade(&objectId, &subjectGrade, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SubjectGradeController) DeleteSubjectGrade(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := controller.SubjectGradeService.DeleteSubjectGrade(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SubjectGradeController) RegisterSubjectGradeRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/subject-grades", controller.CreateSubjectGrade)
	routerGroup.GET("/subject-grades/:id", controller.GetSubjectGrade)
	routerGroup.GET("/subject-grades", controller.GetAllSubjectGrades)
	routerGroup.PATCH("/subject-grades/:id", controller.UpdateSubjectGrade)
	routerGroup.DELETE("/subject-grades/:id", controller.DeleteSubjectGrade)
}
