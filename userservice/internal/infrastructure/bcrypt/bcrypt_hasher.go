package bcrypthash

import (
	"errors"
	"userservice/internal/repository/hasher"

	"golang.org/x/crypto/bcrypt"
)

type BcryptHasher struct{}

func NewBcryptHasher() *BcryptHasher {
	return &BcryptHasher{}
}

func (b *BcryptHasher) Hash(pass []byte) ([]byte, error) {
	hashPass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hashPass, err
}

func (b *BcryptHasher) ComparePassword(hashPass, pass []byte) error {
	err := bcrypt.CompareHashAndPassword(hashPass, pass)
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return hasher.ErrWrongPassword
	}
	return err
}
