package projectdomain

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProjectDomain(t *testing.T) {
	tests := []struct {
		testName  string
		ownerId   uint32
		name      string
		expDomain *ProjectDomain
		expErr    error
	}{
		{
			testName: "Success",
			ownerId:  1,
			name:     "name",
			expDomain: &ProjectDomain{
				OwnerId: 1,
				Name:    "name",
			},
			expErr: nil,
		}, {
			testName:  "Invalid owner id",
			ownerId:   0,
			name:      "name",
			expDomain: nil,
			expErr:    ErrInvalidOwnerId,
		}, {
			testName:  "Invalid name",
			ownerId:   1,
			name:      strings.Repeat("name", 300),
			expDomain: nil,
			expErr:    ErrInvalidName,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			domain, err := NewProjectDomain(tt.ownerId, tt.name)
			require.Equal(t, tt.expErr, err)
			require.Equal(t, tt.expDomain, domain)
		})
	}
}
