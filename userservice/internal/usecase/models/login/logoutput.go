package logmodel

type LoginOutput struct {
	SessionId  string
	FirstName  string
	MiddleName string
	LastName   string
}

func NewLoginOutput(sessionId, firstName, middleName, lastName string) *LoginOutput {
	return &LoginOutput{
		SessionId:  sessionId,
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
	}
}
