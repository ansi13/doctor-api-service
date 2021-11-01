package models

import "gorm.io/gorm"

type Patient struct {
	gorm.Model
	Name        string `json:"name"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email" gorm:"unique"`
	Address     string `json:"address"`
	Observation string `json:"observation"`
	DoctorID    uint   `json:"doctor_id"`
}

type DisplayPatient struct {
	ID          uint   `json:"patient_id"`
	Name        string `json:"name"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Observation string `json:"observation"`
	DoctorName  string `json:"doctor_name"`
}
