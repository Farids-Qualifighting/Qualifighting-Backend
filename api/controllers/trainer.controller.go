package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type TrainerController struct {
	TrainerService services.TrainerService
}

func NewTrainerController(trainerService services.TrainerService) TrainerController {
	return TrainerController{
		TrainerService: trainerService,
	}
}

func (controller *TrainerController) CreateTrainer(ctx *gin.Context) {
	var trainer models.Trainer

	if err := ctx.ShouldBindJSON(&trainer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.TrainerService.CreateTrainer(trainer, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TrainerController) GetTrainer(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	trainer, err := controller.TrainerService.GetTrainer(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, trainer)
}

func (controller *TrainerController) GetAllTrainers(ctx *gin.Context) {
	trainers, err := controller.TrainerService.GetAllTrainers(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, trainers)
}

func (controller *TrainerController) UpdateTrainer(ctx *gin.Context) {
	var trainer models.UpdateTrainer

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&trainer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.TrainerService.UpdateTrainer(&objectId, trainer, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TrainerController) DeleteTrainer(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.TrainerService.DeleteTrainer(&objectId, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *TrainerController) RegisterTrainerRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/trainers", controller.CreateTrainer)
	routerGroup.GET("/trainers/:id", controller.GetTrainer)
	routerGroup.GET("/trainers", controller.GetAllTrainers)
	routerGroup.PATCH("/trainers/:id", controller.UpdateTrainer)
	routerGroup.DELETE("/trainers/:id", controller.DeleteTrainer)
}
