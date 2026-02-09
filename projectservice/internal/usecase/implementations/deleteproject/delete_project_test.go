package deleteproject

import (
	"context"
	"io"
	"log/slog"
	"projectservice/internal/repository/storage"
	deleteerr "projectservice/internal/usecase/error/deleteproject"
	deletemocks "projectservice/internal/usecase/implementations/deleteproject/mocks"
	deletemodel "projectservice/internal/usecase/models/deleteproject"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

//go:generate mockgen -source=./../../../repository/storage/storagerepo.go -destination=./mocks/mock_storage.go -package=deletemocks
func TestDeleteProject(t *testing.T) {
	tests := []struct {
		testName string

		expStorage          bool
		storageInputProjId  uint32
		storageInputOwnerId uint32
		storageReturnErr    error

		deleteInput *deletemodel.DeleteProjectInput

		expErr    error
		expOutput *deletemodel.DeleteProjectOutput
	}{
		{
			testName: "Success",

			expStorage:          true,
			storageInputProjId:  1,
			storageInputOwnerId: 1,
			storageReturnErr:    nil,

			deleteInput: deletemodel.NewDeleteProjectInput(1, 1),

			expErr:    nil,
			expOutput: deletemodel.NewDeleteProjectOutput(true),
		}, {
			testName: "Invalid project id",

			expStorage:          false,
			storageInputProjId:  0,
			storageInputOwnerId: 1,
			storageReturnErr:    storage.ErrNotFound,

			deleteInput: deletemodel.NewDeleteProjectInput(1, 0),

			expErr:    deleteerr.ErrInvalidProjectId,
			expOutput: deletemodel.NewDeleteProjectOutput(false),
		}, {
			testName: "Not found",

			expStorage:          true,
			storageInputProjId:  1,
			storageInputOwnerId: 1,
			storageReturnErr:    storage.ErrNotFound,

			deleteInput: deletemodel.NewDeleteProjectInput(1, 1),

			expErr:    deleteerr.ErrProjectNotFound,
			expOutput: deletemodel.NewDeleteProjectOutput(false),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			log := slog.New(slog.NewTextHandler(io.Discard, nil))

			storageMock := deletemocks.NewMockStorageRepo(ctrl)
			if tt.expStorage {
				storageMock.EXPECT().Delete(gomock.Any(), tt.storageInputOwnerId, tt.storageInputProjId).
					Return(tt.storageReturnErr)
			}

			deleteUC := NewDeleteProjectUC(log, storageMock)

			out, err := deleteUC.Execute(context.Background(), tt.deleteInput)
			assert.Equal(t, tt.expErr, err)
			assert.Equal(t, tt.expOutput, out)
		})
	}
}
