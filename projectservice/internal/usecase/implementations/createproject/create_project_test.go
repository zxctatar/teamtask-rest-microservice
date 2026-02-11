package createproject

import (
	"context"
	"io"
	"log/slog"
	projectdomain "projectservice/internal/domain/project"
	"projectservice/internal/repository/storage"
	createerr "projectservice/internal/usecase/error/createproject"
	createmocks "projectservice/internal/usecase/implementations/createproject/mocks"
	createmodel "projectservice/internal/usecase/models/createproject"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

//go:generate mockgen -source=./../../../repository/storage/storagerepo.go -destination=./mocks/mock_storage.go -package=createmocks
func TestCreateProject(t *testing.T) {
	tests := []struct {
		testName string

		expStorage       bool
		storageInpit     *projectdomain.ProjectDomain
		storageReturnErr error
		storageReturnId  uint32

		createInput *createmodel.CreateProjectInput

		expOut *createmodel.CreateProjectOutput
		expErr error
	}{
		{
			testName: "Success",

			expStorage:       true,
			storageInpit:     &projectdomain.ProjectDomain{OwnerId: 1, Name: "Name"},
			storageReturnErr: nil,
			storageReturnId:  1,

			createInput: createmodel.NewCreateProjectInput(1, "Name"),

			expOut: createmodel.NewCreateProjectOutput(1),
			expErr: nil,
		}, {
			testName: "Invalid domain",

			expStorage:       false,
			storageInpit:     &projectdomain.ProjectDomain{OwnerId: 1, Name: strings.Repeat("Name", 300)},
			storageReturnErr: nil,
			storageReturnId:  1,

			createInput: createmodel.NewCreateProjectInput(1, strings.Repeat("Name", 300)),

			expOut: createmodel.NewCreateProjectOutput(0),
			expErr: projectdomain.ErrInvalidName,
		}, {
			testName: "Already exists",

			expStorage:       true,
			storageInpit:     &projectdomain.ProjectDomain{OwnerId: 1, Name: "Name"},
			storageReturnErr: storage.ErrAlreadyExists,
			storageReturnId:  0,

			createInput: createmodel.NewCreateProjectInput(1, "Name"),

			expOut: createmodel.NewCreateProjectOutput(0),
			expErr: createerr.ErrAlreadyExists,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storageMock := createmocks.NewMockStorageRepo(ctrl)
			if tt.expStorage {
				storageMock.EXPECT().Save(gomock.Any(), tt.storageInpit).
					Return(tt.storageReturnId, tt.storageReturnErr)
			}

			log := slog.New(slog.NewTextHandler(io.Discard, nil))

			createUC := NewCreateProjectUC(log, storageMock)

			out, err := createUC.Execute(context.Background(), tt.createInput)
			require.Equal(t, tt.expErr, err)
			require.Equal(t, tt.expOut, out)
		})
	}
}
