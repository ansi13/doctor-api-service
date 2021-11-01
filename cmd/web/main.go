package main

import (
	"github.com/ansi13/doctor-api-service/cmd/web/doctors"
	"github.com/ansi13/doctor-api-service/cmd/web/patients"
	"github.com/ansi13/doctor-api-service/cmd/web/ping"
	"github.com/ansi13/doctor-api-service/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {

	utils.InitDB()

	router := gin.Default()
	router.GET("/ping", ping.Ping)

	v1 := router.Group("/api/v1/patients")
	{
		v1.GET("/", patients.GetPatients)
		v1.POST("/", patients.CreatePatients)
		v1.GET("/:id", patients.FetchSinglePatient)
		v1.PUT("/:id", patients.UpdateSinglePatient)
		v1.DELETE("/:id", patients.DeleteSinglePatient)
	}

	doctor_v1 := router.Group("/api/v1/doctors")
	{
		doctor_v1.POST("/login", doctors.DoctorLogin)
		doctor_v1.POST("/signup", doctors.CreateDoctor)
	}

	router.Run("localhost:8080")
}
