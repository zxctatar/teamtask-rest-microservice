package authmodel

type AuthOutput struct {
	UserId uint32
}

func NewAuthOutput(userId uint32) *AuthOutput {
	return &AuthOutput{
		UserId: userId,
	}
}
