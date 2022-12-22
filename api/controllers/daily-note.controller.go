package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qualifighting.backend.de/api/services"
	"qualifighting.backend.de/models"
)

type DailyNoteController struct {
	DailyNoteService services.DailyNoteService
}

func NewDailyNoteController(dailyNotesService services.DailyNoteService) DailyNoteController {
	return DailyNoteController{
		DailyNoteService: dailyNotesService,
	}
}

func (controller *DailyNoteController) CreateDailyNote(ctx *gin.Context) {
	var note models.DailyNote

	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	err := controller.DailyNoteService.CreateDailyNote(&note, ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *DailyNoteController) GetDailyNote(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	note, errService := controller.DailyNoteService.GetDailyNote(&objectId, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, note)
}

func (controller *DailyNoteController) GetAllDailyNotes(ctx *gin.Context) {
	notes, err := controller.DailyNoteService.GetAllDailyNotes(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, notes)
}

func (controller *DailyNoteController) UpdateDailyNotes(ctx *gin.Context) {
	var note models.UpdateDailyNote

	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.DailyNoteService.UpdateDailyNote(&objectId, &note, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *DailyNoteController) DeleteDailyNote(ctx *gin.Context) {
	id := ctx.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	errService := controller.DailyNoteService.DeleteDailyNote(&objectId, ctx)

	if errService != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (controller *DailyNoteController) RegisterDailyNoteRoutes(routerGroup *gin.RouterGroup) {

	routerGroup.POST("/daily-notes", controller.CreateDailyNote)
	routerGroup.GET("/daily-notes/:id", controller.GetDailyNote)
	routerGroup.GET("/daily-notes", controller.GetAllDailyNotes)
	routerGroup.PATCH("/daily-notes/:id", controller.UpdateDailyNotes)
	routerGroup.DELETE("/daily-notes/:id", controller.DeleteDailyNote)
}
