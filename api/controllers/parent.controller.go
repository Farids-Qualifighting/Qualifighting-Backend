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
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.ParentService.CreateParent(&parent, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *ParentController) GetParent(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)
	parent, err := controller.ParentService.GetParent(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, parent)
}

func (controller *ParentController) GetAllParents(ctx *gin.Context) {
	parents, err := controller.ParentService.GetAllParents(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, parents)
}

func (controller *ParentController) UpdateParent(ctx *gin.Context) {
	var parent models.UpdateParent

	id := ctx.Param("id")
	objectId, _ := primitive.ObjectIDFromHex(id)

	if err := ctx.ShouldBindJSON(&parent); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	err := controller.ParentService.UpdateParent(&objectId, &parent, ctx)

	if err != nil {
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
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *ParentController) RegisterParentRoutes(routerGroup *gin.RouterGroup) {
	parentRoute := routerGroup.Group("/parent")

	parentRoute.POST("/create", controller.CreateParent)
	parentRoute.GET("/get/:id", controller.GetParent)
	parentRoute.GET("/all", controller.GetAllParents)
	parentRoute.PATCH("/update/:id", controller.UpdateParent)
	parentRoute.DELETE("/delete/:id", controller.DeleteParent)
}
