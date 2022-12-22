package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type SchoolController struct {
	SchoolService services.SchoolService
}

func NewSchoolController(schoolService services.SchoolService) SchoolController {
	return SchoolController{
		SchoolService: schoolService,
	}
}

func (controller *SchoolController) CreateSchool(ctx *gin.Context) {
	var school models.School

	if err := ctx.ShouldBindJSON(&school); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.SchoolService.CreateSchool(&school, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SchoolController) GetSchool(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	school, err := controller.SchoolService.GetSchool(&objectId, ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, school)
}

func (controller *SchoolController) GetAllSchools(ctx *gin.Context) {
	schools, err := controller.SchoolService.GetAllSchools(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, schools)
}

func (controller *SchoolController) UpdateSchool(ctx *gin.Context) {
	var school models.UpdateSchool

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&school); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.SchoolService.UpdateSchool(&objectId, &school, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SchoolController) DeleteSchool(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.SchoolService.DeleteSchool(&objectId, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *SchoolController) RegisterSchoolRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/schools", controller.CreateSchool)
	routerGroup.GET("/schools/:id", controller.GetSchool)
	routerGroup.GET("/schools", controller.GetAllSchools)
	routerGroup.PATCH("/schools/:id", controller.UpdateSchool)
	routerGroup.DELETE("/schools/:id", controller.DeleteSchool)
}
