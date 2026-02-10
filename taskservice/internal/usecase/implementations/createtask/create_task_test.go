package createuc

import (
	"context"
	"io"
	"log/slog"
	"strings"
	taskdomain "taskservice/internal/domain/task"
	createmocks "taskservice/internal/usecase/implementations/createtask/mocks"
	createmodel "taskservice/internal/usecase/models/createtask"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

//go:generate mockgen -source=./../../../repository/storage/storagerepo.go -destination=./mocks/mock_storage.go -package=createmocks
func TestCreateUC(t *testing.T) {
	timeNow := time.Now()

	tests := []struct {
		testName string

		expStorage    bool
		storInput     *taskdomain.TaskDomain
		storReturn    uint32
		storReturnErr error

		in     *createmodel.CreateTaskInput
		expOut *createmodel.CreateTaskOutput
		expErr error
	}{
		{
			testName: "Success",

			expStorage: true,
			storInput: &taskdomain.TaskDomain{
				Id:          0,
				ProjectId:   1,
				Description: "desc",
				Deadline:    timeNow,
			},
			storReturn:    1,
			storReturnErr: nil,

			in: createmodel.NewCreateInput(
				1,
				"desc",
				timeNow,
			),
			expOut: createmodel.NewCreateOutput(
				1,
			),
			expErr: nil,
		}, {
			testName: "Bad project id",

			expStorage: false,
			storInput: &taskdomain.TaskDomain{
				Id:          0,
				ProjectId:   1,
				Description: "desc",
				Deadline:    timeNow,
			},
			storReturn:    1,
			storReturnErr: nil,

			in: createmodel.NewCreateInput(
				0,
				"desc",
				timeNow,
			),
			expOut: nil,
			expErr: taskdomain.ErrInvalidProjectId,
		}, {
			testName: "Bad description",

			expStorage: false,
			storInput: &taskdomain.TaskDomain{
				Id:          0,
				ProjectId:   1,
				Description: "desc",
				Deadline:    timeNow,
			},
			storReturn:    1,
			storReturnErr: nil,

			in: createmodel.NewCreateInput(
				1,
				strings.Repeat("A", 256),
				timeNow,
			),
			expOut: nil,
			expErr: taskdomain.ErrInvalidDescription,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			storMock := createmocks.NewMockStorageRepo(ctrl)
			if tt.expStorage {
				storMock.EXPECT().Save(gomock.Any(), tt.storInput).
					Return(tt.storReturn, tt.storReturnErr)
			}

			log := slog.New(slog.NewTextHandler(io.Discard, nil))

			createUC := NewCreateUC(log, storMock)

			out, err := createUC.Execute(context.Background(), tt.in)
			require.Equal(t, tt.expErr, err)
			require.Equal(t, tt.expOut, out)
		})
	}
}
