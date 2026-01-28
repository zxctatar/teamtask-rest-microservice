package storage

import "context"

type StorageRepo interface {
	Save(ctx context.Context, ownerId uint32, name string) error
}
