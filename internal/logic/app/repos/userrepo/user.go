package userrepo

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/user"
	"github.com/google/uuid"
)

type UserStore interface {
	Create(ctx context.Context, u user.User) (*uuid.UUID, error)
	Read(ctx context.Context, uid uuid.UUID) (*user.User, error)
	Delete(ctx context.Context, uid uuid.UUID) error
	SearchUsers(ctx context.Context, s string) (chan user.User, error)
}

type Users struct {
	ustore UserStore
}

func New(ustore UserStore) *Users {
	return &Users{
		ustore: ustore,
	}
}

func (us *Users) Create(ctx context.Context, u user.User) (*uuid.UUID, error) {
	return nil, nil
}

func (us *Users) Read(ctx context.Context, uid uuid.UUID) (*user.User, error) {
	return nil, nil
}

func (us *Users) Delete(ctx context.Context, uid uuid.UUID) error {
	return nil
}

func (us *Users) SearchUsers(ctx context.Context, s string) (chan user.User, error) {
	return nil, nil
}
