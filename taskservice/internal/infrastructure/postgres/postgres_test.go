package postgres

import (
	"context"
	"regexp"
	taskdomain "taskservice/internal/domain/task"
	posmodels "taskservice/internal/infrastructure/postgres/models"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestPostgres_Save_Success(t *testing.T) {
	timeNow := time.Now()

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	posModel := posmodels.NewTaskPosModel(
		0,
		1,
		"desc",
		timeNow,
	)

	mock.ExpectQuery(regexp.QuoteMeta(QuerieCreate)).
		WithArgs(posModel.ProjectId, posModel.Description, posModel.Deadline).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)).
		WillReturnError(nil)

	postgres := NewPostgres(db)

	td, err := taskdomain.NewTaskDomain(
		1,
		"desc",
		timeNow,
	)
	require.NoError(t, err)

	id, err := postgres.Save(context.Background(), td)
	require.NoError(t, err)
	require.Equal(t, uint32(1), id)
}
