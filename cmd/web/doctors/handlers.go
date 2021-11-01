package doctors

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/ansi13/doctor-api-service/pkg/forms"
	"github.com/ansi13/doctor-api-service/pkg/models"
	"github.com/ansi13/doctor-api-service/pkg/utils"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func DoctorLogin(c *gin.Context) {
	var form_data forms.LoginForm
	var doctor_data models.Doctor

	if err := c.Bind(&form_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result := utils.DB.First(&doctor_data, "username = ?", form_data.Username)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not found"})
		return
	}

	if !utils.CheckPasswordHash(form_data.Password, doctor_data.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "password incorrect"})
		return
	}

	token, err := utils.GenerateToken(form_data.Username, []byte("hello@dude"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "token generation failed"})
	}

	c.JSON(http.StatusOK, gin.H{"Authorized": fmt.Sprintf("Bearer %s", token)})
}

func CreateDoctor(c *gin.Context) {
	var form_data forms.SignUpForm

	if err := c.Bind(&form_data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	hashed_password, _ := utils.HashPassword(form_data.Password)

	doctor := models.Doctor{Name: form_data.Name,
		Email: form_data.Email, Username: form_data.Username, Password: hashed_password}

	result := utils.DB.Create(&doctor)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "DB insertion error"})
		return
	}

	log.Printf("new user created: %s", form_data.Username)
	c.JSON(http.StatusCreated, gin.H{"id": doctor.ID, "username": doctor.Username, "created_at": doctor.CreatedAt})
}
