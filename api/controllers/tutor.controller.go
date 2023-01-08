package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type TutorController struct {
	TutorService services.TutorService
}

func NewTutorController(tutorService services.TutorService) TutorController {
	return TutorController{
		TutorService: tutorService,
	}
}

func (controller *TutorController) CreateTutor(ctx *gin.Context) {
	var tutor models.Tutor

	if err := ctx.ShouldBindJSON(&tutor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.TutorService.CreateTutor(tutor, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TutorController) GetTutor(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	tutor, err := controller.TutorService.GetTutor(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tutor)
}

func (controller *TutorController) GetAllTutors(ctx *gin.Context) {
	tutors, err := controller.TutorService.GetAllTutors(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tutors)
}

func (controller *TutorController) UpdateTutor(ctx *gin.Context) {
	var tutor models.UpdateTutor

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&tutor); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.TutorService.UpdateTutor(&objectId, tutor, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TutorController) DeleteTutor(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.TutorService.DeleteTutor(&objectId, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TutorController) RegisterTutorRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/tutors", controller.CreateTutor)
	routerGroup.GET("/tutors/:id", controller.GetTutor)
	routerGroup.GET("/tutors", controller.GetAllTutors)
	routerGroup.PATCH("/tutors/:id", controller.UpdateTutor)
	routerGroup.DELETE("/tutors/:id", controller.DeleteTutor)
}
