package groupmemstore

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/group"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/grouprepo"
	"github.com/google/uuid"
	"sync"
)

var _ grouprepo.GroupStore = &Groups{}

type Groups struct {
	sync.Mutex
	m map[uuid.UUID]group.Group
}

func New() *Groups {
	return &Groups{
		m: make(map[uuid.UUID]group.Group),
	}
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
