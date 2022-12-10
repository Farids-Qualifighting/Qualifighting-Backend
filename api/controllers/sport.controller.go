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
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.SportService.CreateSport(&sport, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SportController) GetSport(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	student, err := controller.SportService.GetSport(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}

func (controller *SportController) GetAllSports(ctx *gin.Context) {
	sports, err := controller.SportService.GetAllSports(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, sports)
}

func (controller *SportController) UpdateSport(ctx *gin.Context) {
	var sport models.UpdateSport

	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	if err := ctx.ShouldBindJSON(&sport); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.SportService.UpdateSport(&objectId, &sport, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SportController) DeleteSport(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := controller.SportService.DeleteSport(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SportController) RegisterSportRoutes(routerGroup *gin.RouterGroup) {
	sportRoute := routerGroup.Group("/sport")

	sportRoute.POST("/create", controller.CreateSport)
	sportRoute.GET("/get/:id", controller.GetSport)
	sportRoute.GET("/all", controller.GetAllSports)
	sportRoute.PATCH("/update/:id", controller.UpdateSport)
	sportRoute.DELETE("/delete/:id", controller.DeleteSport)
}
