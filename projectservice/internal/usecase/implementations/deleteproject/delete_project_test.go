package deleteproject

import (
	"context"
	"io"
	"log/slog"
	projectdomain "projectservice/internal/domain/project"
	"projectservice/internal/repository/storage"
	deleteerr "projectservice/internal/usecase/error/deleteproject"
	deletemocks "projectservice/internal/usecase/implementations/deleteproject/mocks"
	deletemodel "projectservice/internal/usecase/models/deleteproject"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

//go:generate mockgen -source=./../../../repository/storage/storage.go -destination=./mocks/mock_storage.go -package=deletemocks
func TestDeleteProject(t *testing.T) {
	tests := []struct {
		testName string

		expStorage       bool
		storageInput     *projectdomain.ProjectDomain
		storageReturnErr error

		deleteInput *deletemodel.DeleteProjectInput

		expErr    error
		expOutput *deletemodel.DeleteProjectOutput
	}{
		{
			testName: "Success",

			expStorage:       true,
			storageInput:     &projectdomain.ProjectDomain{OwnerId: 1, Name: "Name"},
			storageReturnErr: nil,

			deleteInput: deletemodel.NewDeleteProjectInput(1, "Name"),

			expErr:    nil,
			expOutput: deletemodel.NewDeleteProjectOutput(true),
		}, {
			testName: "Invalid domain",

			expStorage:       false,
			storageInput:     &projectdomain.ProjectDomain{OwnerId: 1, Name: strings.Repeat("Name", 300)},
			storageReturnErr: nil,

			deleteInput: deletemodel.NewDeleteProjectInput(1, strings.Repeat("Name", 300)),

			expErr:    projectdomain.ErrInvalidName,
			expOutput: deletemodel.NewDeleteProjectOutput(false),
		}, {
			testName: "Not found",

			expStorage:       true,
			storageInput:     &projectdomain.ProjectDomain{OwnerId: 1, Name: "Name"},
			storageReturnErr: storage.ErrNotFound,

			deleteInput: deletemodel.NewDeleteProjectInput(1, "Name"),

			expErr:    deleteerr.ErrProjectNotFound,
			expOutput: deletemodel.NewDeleteProjectOutput(false),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			log := slog.New(slog.NewTextHandler(io.Discard, nil))

			storageMock := deletemocks.NewMockStorage(ctrl)
			if tt.expStorage {
				storageMock.EXPECT().Delete(gomock.Any(), tt.storageInput).
					Return(tt.storageReturnErr)
			}

			deleteUC := NewDeleteProjectUC(log, storageMock)

			out, err := deleteUC.Execute(context.Background(), tt.deleteInput)
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expOutput, out)
		})
	}
}
