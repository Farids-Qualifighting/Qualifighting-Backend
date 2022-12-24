package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type ParentController struct {
	ParentService services.ParentService
}

func NewParentController(parentService services.ParentService) ParentController {
	return ParentController{
		ParentService: parentService,
	}
}

func (controller *ParentController) CreateParent(ctx *gin.Context) {
	var parent models.Parent

	if err := ctx.ShouldBindJSON(&parent); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.ParentService.CreateParent(&parent, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *ParentController) GetParent(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	parent, err := controller.ParentService.GetParent(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, parent)
}

func (controller *ParentController) GetAllParents(ctx *gin.Context) {
	parents, err := controller.ParentService.GetAllParents(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, parents)
}

// FIXME Address gets overwritten when updating other properties
func (controller *ParentController) UpdateParent(ctx *gin.Context) {
	var parent models.UpdateParent

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&parent); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.ParentService.UpdateParent(&objectId, &parent, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *ParentController) DeleteParent(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	err := controller.ParentService.DeleteParent(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *ParentController) RegisterParentRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/parents", controller.CreateParent)
	routerGroup.GET("/parents/:id", controller.GetParent)
	routerGroup.GET("/parents", controller.GetAllParents)
	routerGroup.PATCH("/parents/:id", controller.UpdateParent)
	routerGroup.DELETE("/parents/:id", controller.DeleteParent)
}
