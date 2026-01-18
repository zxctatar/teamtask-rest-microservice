package uuidgen

import "github.com/google/uuid"

type UUIDGenerator struct{}

func NewUUIDGenerator() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (u *UUIDGenerator) New() string {
	return uuid.NewString()
}
