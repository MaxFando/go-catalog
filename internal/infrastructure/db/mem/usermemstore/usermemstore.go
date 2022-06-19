package usermemstore

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/user"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/userrepo"
	"github.com/google/uuid"
	"sync"
)

var _ userrepo.UserStore = &Users{}

type Users struct {
	sync.Mutex
	m map[uuid.UUID]user.User
}

func New() *Users {
	return &Users{
		m: make(map[uuid.UUID]user.User),
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

func (us *Users) Columns() string {
	return "id, name, data, permissions"
}

func (us *Users) Fields(u *user.User) []interface{} {
	return []interface{}{&u.ID, &u.Name, &u.Data, &u.Permissions}
}
