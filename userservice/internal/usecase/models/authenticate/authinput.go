package authmodel

type AuthInput struct {
	SessionId string
}

func NewAuthInput(sessionId string) *AuthInput {
	return &AuthInput{
		SessionId: sessionId,
	}
}
