package grouprepo

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/group"
	"github.com/google/uuid"
)

type GroupStore interface {
	Create(ctx context.Context, cg group.Group) (*uuid.UUID, error)
	Read(ctx context.Context, cgId uuid.UUID) (*uuid.UUID, error)
	Delete(ctx context.Context, cgId uuid.UUID) error
	Search(ctx context.Context, s string) (chan group.Group, error)
	SearchByUsersName(ctx context.Context, names []string) (chan group.Group, error)

	AttachUser(ctx context.Context, cg group.Group, uid uuid.UUID) error
	DetachUser(ctx context.Context, cg group.Group, uid uuid.UUID) error
}

type Groups struct {
	store GroupStore
}

func (g *Groups) Create(ctx context.Context, cg group.Group) (*uuid.UUID, error) {
	return nil, nil
}

func (g *Groups) Read(ctx context.Context, cgId uuid.UUID) (*uuid.UUID, error) {
	return nil, nil
}

func (g *Groups) Delete(ctx context.Context, cgId uuid.UUID) error {
	return nil
}

func (g *Groups) Search(ctx context.Context, s string) (chan group.Group, error) {
	return nil, nil
}

func (g *Groups) SearchByUsersName(ctx context.Context, names []string) (chan group.Group, error) {
	return nil, nil
}

func (g *Groups) AttachUser(ctx context.Context, cg group.Group, uid uuid.UUID) error {
	return nil
}

func (g *Groups) DetachUser(ctx context.Context, cg group.Group, uid uuid.UUID) error {
	return nil
}

func New(store GroupStore) *Groups {
	return &Groups{
		store: store,
	}
}
