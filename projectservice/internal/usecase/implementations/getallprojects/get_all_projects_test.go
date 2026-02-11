package getallprojects

import (
	"context"
	"io"
	"log/slog"
	projectdomain "projectservice/internal/domain/project"
	"projectservice/internal/repository/storage"
	getallerr "projectservice/internal/usecase/error/getallprojects"
	getallmocks "projectservice/internal/usecase/implementations/getallprojects/mocks"
	getallmodel "projectservice/internal/usecase/models/getallprojects"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

//go:generate mockgen -source=./../../../repository/storage/storagerepo.go -destination=./mocks/mock_storage.go -package=getallmocks
func TestGetAllProjects(t *testing.T) {
	timeNow := time.Now()

	tests := []struct {
		testName string

		storageInput     uint32
		storageOutput    []*projectdomain.ProjectDomain
		storageReturnErr error

		in *getallmodel.GetAllProjectsInput

		expErr error
		expOut *getallmodel.GetAllProjectsOutput
	}{
		{
			testName: "Success",

			storageInput: 1,
			storageOutput: []*projectdomain.ProjectDomain{
				&projectdomain.ProjectDomain{Id: 1, OwnerId: 1, Name: "A", CreatedAt: timeNow},
			},
			storageReturnErr: nil,

			in: getallmodel.NewGetAllProjectsInput(1),

			expErr: nil,
			expOut: &getallmodel.GetAllProjectsOutput{Projects: []*projectdomain.ProjectDomain{
				&projectdomain.ProjectDomain{Id: 1, OwnerId: 1, Name: "A", CreatedAt: timeNow},
			}},
		}, {
			testName: "Not found",

			storageInput:     1,
			storageOutput:    nil,
			storageReturnErr: storage.ErrNotFound,

			in: getallmodel.NewGetAllProjectsInput(1),

			expErr: getallerr.ErrProjectsNotFound,
			expOut: &getallmodel.GetAllProjectsOutput{Projects: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			log := slog.New(slog.NewTextHandler(io.Discard, nil))

			storageMock := getallmocks.NewMockStorageRepo(ctrl)
			storageMock.EXPECT().GetAll(gomock.Any(), tt.storageInput).
				Return(tt.storageOutput, tt.storageReturnErr)

			handl := NewGetAllProjectsUC(log, storageMock)

			out, err := handl.Execute(context.Background(), tt.in)
			require.Equal(t, tt.expErr, err)
			require.Equal(t, tt.expOut, out)
		})
	}
}
