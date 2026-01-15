package regdto

type RegistrationRequest struct {
	FirstName  string `json:"first_name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	MiddleName string `json:"middle_name"`
	Age        uint32 `json:"age" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
}
