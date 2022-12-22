package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type SportController struct {
	SportService services.SportService
}

func NewSportController(sportService services.SportService) SportController {
	return SportController{
		SportService: sportService,
	}
}

func (controller *SportController) CreateSport(ctx *gin.Context) {
	var sport models.Sport

	if err := ctx.ShouldBindJSON(&sport); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.SportService.CreateSport(&sport, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SportController) GetSport(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	student, err := controller.SportService.GetSport(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (controller *SportController) GetAllSports(ctx *gin.Context) {
	sports, err := controller.SportService.GetAllSports(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sports)
}

func (controller *SportController) UpdateSport(ctx *gin.Context) {
	var sport models.UpdateSport

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&sport); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.SportService.UpdateSport(&objectId, &sport, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SportController) DeleteSport(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.SportService.DeleteSport(&objectId, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SportController) RegisterSportRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/sports", controller.CreateSport)
	routerGroup.GET("/sports/:id", controller.GetSport)
	routerGroup.GET("/sports", controller.GetAllSports)
	routerGroup.PATCH("/sports/:id", controller.UpdateSport)
	routerGroup.DELETE("/sports/:id", controller.DeleteSport)
}
