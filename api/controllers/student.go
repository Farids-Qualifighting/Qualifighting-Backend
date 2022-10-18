package controllers

import (
	"context"
	"net/http"

	"firebase.google.com/go/auth"
	"firebase.google.com/go/v4/errorutils"
	"github.com/gin-gonic/gin"

	"qualifighting.backend.de/lib"

	"qualifighting.backend.de/models"
)

type User struct{}

func (User) Create(c *gin.Context) {
	var request *models.CreateStudentData
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user := c.MustGet("user").(*models.JWTUserMiddleware)

	// only super admins can create stuedents with different school ids
	if user.Level < int(models.UserStatusSuperAdmin) {
		c.JSON(http.StatusForbidden, gin.H{"error": "FORBIDDEN"})
		return
	}

	mongo := lib.MongoDBSchoolCollection()
	firebase := lib.GetFirebaseAuth()

	params := (&auth.UserToCreate{}).
		Email(request.Email).
		EmailVerified(false).
		Password(request.Password).
		DisplayName(request.FirstName + " " + request.LastName)

	userRecord, err := firebase.CreateUser(c, params)
	if err != nil {
		if errorutils.IsAlreadyExists(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "STUDENT_ALREADY_EXISTS"})
			return
		}
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = firebase.SetCustomUserClaims(c, userRecord.UID, gin.H{"level": models.UserStatusStudent, "school_id": request.SchoolID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	userData := models.StudentData{
		ID:                 userRecord.UID,
		Level:              uint8(models.UserStatusStudent),
		FirstName:          request.FirstName,
		LastName:           request.LastName,
		Address:            request.Address,
		SchoolID:           request.SchoolID,
		BirthDate:          request.BirthDate,
		Phone:              request.Phone,
		Gender:             request.Gender,
		Email:              request.Email,
		ClassTeacher:       request.ClassTeacher,
		Certificates:       make([]string, 0),
		SocialMediaContact: request.SocialMediaContact,
	}

	_, err = mongo.InsertOne(context.TODO(), userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "CUSTOMER_CREATED"})
}
