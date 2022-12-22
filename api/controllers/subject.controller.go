package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type SubjectController struct {
	SubjectService services.SubjectService
}

func NewSubjectController(subjectService services.SubjectService) SubjectController {
	return SubjectController{
		SubjectService: subjectService,
	}
}

func (controller *SubjectController) CreateSubject(ctx *gin.Context) {
	var subject models.Subject

	if err := ctx.ShouldBindJSON(&subject); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.SubjectService.CreateSubject(&subject, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SubjectController) GetSubject(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	student, err := controller.SubjectService.GetSubject(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (controller *SubjectController) GetAllSubjects(ctx *gin.Context) {
	subjects, err := controller.SubjectService.GetAllSubjects(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, subjects)
}

func (controller *SubjectController) UpdateSubject(ctx *gin.Context) {
	var subject models.UpdateSubject

	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	if err := ctx.ShouldBindJSON(&subject); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.SubjectService.UpdateSubject(&objectId, &subject, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SubjectController) DeleteSubject(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := controller.SubjectService.DeleteSubject(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SubjectController) RegisterSubjectRoutes(routerGroup *gin.RouterGroup) {
	subjectRoute := routerGroup.Group("/subject")

	subjectRoute.POST("/create", controller.CreateSubject)
	subjectRoute.GET("/get/:id", controller.GetSubject)
	subjectRoute.GET("/all", controller.GetAllSubjects)
	subjectRoute.PATCH("/update/:id", controller.UpdateSubject)
	subjectRoute.DELETE("/delete/:id", controller.DeleteSubject)
}
