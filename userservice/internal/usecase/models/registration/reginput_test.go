package regmodel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegInput(t *testing.T) {
	tests := []struct {
		testName   string
		FirstName  string
		MiddleName string
		LastName   string
		Password   string
		Email      string
		ExpError   error
	}{
		{
			testName:   "Success",
			FirstName:  "Ivan",
			MiddleName: "Ivanov",
			LastName:   "Ivanovich",
			Password:   "somePass",
			Email:      "gmail@gmail.com",
			ExpError:   nil,
		},
		{
			testName:   "Invalid first name",
			FirstName:  "",
			MiddleName: "Ivanov",
			LastName:   "Ivanovich",
			Password:   "somePass",
			Email:      "gmail@gmail.com",
			ExpError:   ErrEmptyFirstName,
		},
		{
			testName:   "Without middle name",
			FirstName:  "Ivan",
			MiddleName: "",
			LastName:   "Ivanovich",
			Password:   "somePass",
			Email:      "gmail@gmail.com",
			ExpError:   nil,
		},
		{
			testName:   "Invalid last name",
			FirstName:  "Ivan",
			MiddleName: "Ivanov",
			LastName:   "",
			Password:   "somePass",
			Email:      "gmail@gmail.com",
			ExpError:   ErrEmptyLastName,
		},
		{
			testName:   "Invalid password",
			FirstName:  "Ivan",
			MiddleName: "Ivanov",
			LastName:   "Ivanovich",
			Password:   "",
			Email:      "gmail@gmail.com",
			ExpError:   ErrEmptyPassword,
		},
		{
			testName:   "Invalid email",
			FirstName:  "Ivan",
			MiddleName: "Ivanov",
			LastName:   "Ivanovich",
			Password:   "somePass",
			Email:      "",
			ExpError:   ErrEmptyEmail,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			rIn := NewRegInput(
				tt.FirstName,
				tt.MiddleName,
				tt.LastName,
				tt.Password,
				tt.Email,
			)
			assert.Equal(t, tt.FirstName, rIn.FirstName)
			assert.Equal(t, tt.MiddleName, rIn.MiddleName)
			assert.Equal(t, tt.LastName, rIn.LastName)
			assert.Equal(t, tt.Password, rIn.Password)
			assert.Equal(t, tt.Email, rIn.Email)
		})
	}
}
