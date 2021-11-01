package forms

type CreatePatientForm struct {
	Name        string `json:"name" binding:"required"`
	Age         string `json:"age" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Address     string `json:"address" binding:"required"`
	Observation string `json:"observation"`
}

type UpdatePatientForm struct {
	Name        string `json:"name"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Observation string `json:"observation"`
}
