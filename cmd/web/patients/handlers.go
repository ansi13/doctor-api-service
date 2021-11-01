package patients

import (
	"errors"
	"net/http"

	"github.com/ansi13/doctor-api-service/pkg/forms"
	"github.com/ansi13/doctor-api-service/pkg/models"
	"github.com/ansi13/doctor-api-service/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPatients(c *gin.Context) {
	username, err := utils.ValidateToken(c.Request, []byte("hello@dude"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Authentication header validation failed"})
		return
	}
	if username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "Authentication header validation failed"})
		return
	}

	var doctor models.Doctor
	result := utils.DB.First(&doctor, "username = ?", username)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Doctor not found"})
		return
	}

	var _patients []models.Patient
	var patients []models.DisplayPatient
	utils.DB.Find(&_patients, "doctor_id = ?", doctor.ID)
	for _, item := range _patients {
		patients = append(patients, models.DisplayPatient{ID: item.ID,
			Name: item.Name, Age: item.Age, Gender: item.Gender, Email: item.Email,
			Observation: item.Observation, Address: item.Address, DoctorName: doctor.Name})
	}

	c.JSON(http.StatusOK, gin.H{"data": patients})
}

func CreatePatients(c *gin.Context) {
	var newPatient forms.CreatePatientForm
	if err := c.Bind(&newPatient); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	username, err := utils.ValidateToken(c.Request, []byte("hello@dude"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Authentication header validation failed"})
		return
	}
	if username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "Authentication header validation failed"})
		return
	}

	var doctor models.Doctor
	result := utils.DB.First(&doctor, "username = ?", username)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Dcotor not found"})
		return
	}

	patient := models.Patient{
		Name:        newPatient.Name,
		Age:         newPatient.Age,
		Gender:      newPatient.Gender,
		Email:       newPatient.Email,
		Address:     newPatient.Address,
		Observation: newPatient.Observation,
		DoctorID:    doctor.ID,
	}

	result = utils.DB.Create(&patient)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "DB insertion error"})
		return
	}

	c.JSON(http.StatusCreated, models.DisplayPatient{ID: patient.ID,
		Name: patient.Name, Age: patient.Age, Gender: patient.Gender, Address: patient.Address,
		Email: patient.Email, Observation: patient.Observation, DoctorName: doctor.Name})
}

func FetchSinglePatient(c *gin.Context) {
	username, err := utils.ValidateToken(c.Request, []byte("hello@dude"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Authentication header validation failed"})
		return
	}
	if username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "Authentication header validation failed"})
		return
	}

	var doctor models.Doctor
	result := utils.DB.First(&doctor, "username = ?", username)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Doctor not found"})
		return
	}
	patient_id := c.Param("id")
	var _patients models.Patient

	result = utils.DB.First(&_patients, "id = ? AND doctor_id = ?", patient_id, doctor.ID)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Patient not found"})
		return
	}

	c.JSON(http.StatusOK, &models.DisplayPatient{ID: _patients.ID,
		Name: _patients.Name, Age: _patients.Age, Gender: _patients.Gender, Email: _patients.Email,
		Address: _patients.Address, Observation: _patients.Observation, DoctorName: doctor.Name})
}

func UpdateSinglePatient(c *gin.Context) {
	var patientForm forms.UpdatePatientForm
	if err := c.Bind(&patientForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	username, err := utils.ValidateToken(c.Request, []byte("hello@dude"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Authentication header validation failed"})
		return
	}
	if username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "Authentication header validation failed"})
		return
	}

	var doctor models.Doctor
	result := utils.DB.First(&doctor, "username = ?", username)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Dcotor not found"})
		return
	}

	patient_id := c.Param("id")

	var _patient models.Patient
	result = utils.DB.First(&_patient, "id = ? AND doctor_id = ?", patient_id, doctor.ID)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Patient not found"})
		return
	}

	if patientForm.Name != "" {
		_patient.Name = patientForm.Name
	}
	if patientForm.Age != "" {
		_patient.Age = patientForm.Age
	}
	if patientForm.Address != "" {
		_patient.Address = patientForm.Address
	}
	if patientForm.Email != "" {
		_patient.Email = patientForm.Email
	}
	if patientForm.Gender != "" {
		_patient.Gender = patientForm.Gender
	}
	if patientForm.Observation != "" {
		_patient.Observation = patientForm.Observation
	}

	utils.DB.Save(&_patient)

	c.JSON(http.StatusOK, models.DisplayPatient{ID: _patient.ID, Name: _patient.Name,
		Age: _patient.Age, Gender: _patient.Gender, Address: _patient.Address,
		Email: _patient.Email, Observation: _patient.Observation, DoctorName: doctor.Name})
}

func DeleteSinglePatient(c *gin.Context) {
	username, err := utils.ValidateToken(c.Request, []byte("hello@dude"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "Authentication header validation failed"})
		return
	}
	if username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"message": "Authentication header validation failed"})
		return
	}

	var doctor models.Doctor
	result := utils.DB.First(&doctor, "username = ?", username)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Dcotor not found"})
		return
	}

	patient_id := c.Param("id")

	var _patient models.Patient
	result = utils.DB.First(&_patient, "id = ? AND doctor_id = ?", patient_id, doctor.ID)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Patient not found"})
		return
	}

	utils.DB.Unscoped().Delete(&_patient)
	c.String(http.StatusNoContent, "")
}
