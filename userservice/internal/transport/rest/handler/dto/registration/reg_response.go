package regdto

type RegistrationResponse struct {
	IsRegistered bool `json:"registered" binding:"required"`
}
