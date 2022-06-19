package communityrepo

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/community"
	"github.com/google/uuid"
)

type CommunityStore interface {
	Create(ctx context.Context, co community.Community) (*uuid.UUID, error)
	Read(ctx context.Context, coId uuid.UUID) (*uuid.UUID, error)
	Delete(ctx context.Context, coId uuid.UUID) error
	Search(ctx context.Context, s string) (chan community.Community, error)
	SearchByUsersName(ctx context.Context, names []string) (chan community.Community, error)

	AttachUser(ctx context.Context, co community.Community, uid uuid.UUID) error
	DetachUser(ctx context.Context, co community.Community, uid uuid.UUID) error
}

type Communities struct {
	store CommunityStore
}

func New(store CommunityStore) *Communities {
	return &Communities{
		store: store,
	}
}

func (c *Communities) Create(ctx context.Context, co community.Community) (*uuid.UUID, error) {
	return nil, nil
}

func (c *Communities) Read(ctx context.Context, coId uuid.UUID) (*uuid.UUID, error) {
	return nil, nil
}

func (c *Communities) Delete(ctx context.Context, coId uuid.UUID) error {
	return nil
}

func (c *Communities) Search(ctx context.Context, co string) (chan community.Community, error) {
	return nil, nil
}

func (c *Communities) SearchByUsersName(ctx context.Context, names []string) (chan community.Community, error) {
	return nil, nil
}

func (c *Communities) AttachUser(ctx context.Context, co community.Community, uid uuid.UUID) error {
	return nil
}

func (c *Communities) DetachUser(ctx context.Context, co community.Community, uid uuid.UUID) error {
	return nil
}
