package communiitymemstore

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/community"
	"github.com/google/uuid"
	"sync"
)

type Communities struct {
	sync.Mutex
	m map[uuid.UUID]community.Community
}

func New() *Communities {
	return &Communities{
		m: make(map[uuid.UUID]community.Community),
	}
}

func (ss *Communities) Create(ctx context.Context, o community.Community) (*uuid.UUID, error) {
	return nil, nil
}

func (ss *Communities) Read(ctx context.Context, oId uuid.UUID) (*uuid.UUID, error) {
	return nil, nil
}

func (ss *Communities) Delete(ctx context.Context, oId uuid.UUID) error {
	return nil
}

func (ss *Communities) Search(ctx context.Context, s string) (chan community.Community, error) {
	return nil, nil
}

func (ss *Communities) SearchByUsersName(ctx context.Context, names []string) (chan community.Community, error) {
	return nil, nil
}

func (ss *Communities) AttachUser(ctx context.Context, o community.Community, uid uuid.UUID) error {
	return nil
}

func (ss *Communities) DetachUser(ctx context.Context, o community.Community, uid uuid.UUID) error {
	return nil
}
