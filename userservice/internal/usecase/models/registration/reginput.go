package regmodel

type RegInput struct {
	FirstName  string
	MiddleName string
	LastName   string
	Password   string
	Email      string
}

func NewRegInput(firstName, middleName, lastName, password, email string) *RegInput {
	return &RegInput{
		FirstName:  firstName,
		MiddleName: middleName,
		LastName:   lastName,
		Password:   password,
		Email:      email,
	}
}
