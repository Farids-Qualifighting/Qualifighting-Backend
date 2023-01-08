package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type CompetitionController struct {
	CompetitionService services.CompetitionService
}

func NewCompetitionController(competitionService services.CompetitionService) CompetitionController {
	return CompetitionController{
		CompetitionService: competitionService,
	}
}

func (controller *CompetitionController) CreateCompetition(ctx *gin.Context) {
	var competition models.Competition

	if err := ctx.ShouldBindJSON(&competition); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.CompetitionService.CreateCompetition(competition, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *CompetitionController) GetCompetition(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	competition, err := controller.CompetitionService.GetCompetition(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, competition)
}

func (controller *CompetitionController) GetAllCompetitions(ctx *gin.Context) {
	competitions, err := controller.CompetitionService.GetAllCompetitions(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, competitions)
}

func (controller *CompetitionController) UpdateCompetition(ctx *gin.Context) {
	var competition models.UpdateCompetition

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&competition); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.CompetitionService.UpdateCompetition(&objectId, &competition, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *CompetitionController) DeleteCompetition(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.CompetitionService.DeleteCompetition(&objectId, ctx)

	if errService != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *CompetitionController) RegisterCompetitionRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/competitions", controller.CreateCompetition)
	routerGroup.GET("/competitions/:id", controller.GetCompetition)
	routerGroup.GET("/competitions", controller.GetAllCompetitions)
	routerGroup.PATCH("/competitions/:id", controller.UpdateCompetition)
	routerGroup.DELETE("/competitions/:id", controller.DeleteCompetition)
}
