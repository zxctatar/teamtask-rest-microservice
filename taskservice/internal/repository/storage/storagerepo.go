package storage

import (
	"context"
	taskdomain "taskservice/internal/domain/task"
)

type StorageRepo interface {
	Save(ctx context.Context, td *taskdomain.TaskDomain) (uint32, error)
}
