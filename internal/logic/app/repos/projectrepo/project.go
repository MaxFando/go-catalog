package projectrepo

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/project"
	"github.com/google/uuid"
)

type ProjectStore interface {
	Create(ctx context.Context, o project.Project) (*uuid.UUID, error)
	Read(ctx context.Context, oId uuid.UUID) (*uuid.UUID, error)
	Delete(ctx context.Context, oId uuid.UUID) error
	Search(ctx context.Context, s string) (chan project.Project, error)
	SearchByUsersName(ctx context.Context, names []string) (chan project.Project, error)

	AttachUser(ctx context.Context, o project.Project, uid uuid.UUID) error
	DetachUser(ctx context.Context, o project.Project, uid uuid.UUID) error
}

type Projects struct {
	store ProjectStore
}

func New(store ProjectStore) *Projects {
	return &Projects{
		store: store,
	}
}

func (p *Projects) Create(ctx context.Context, o project.Project) (*uuid.UUID, error) {
	return nil, nil
}

func (p *Projects) Read(ctx context.Context, oId uuid.UUID) (*uuid.UUID, error) {
	return nil, nil
}

func (p *Projects) Delete(ctx context.Context, oId uuid.UUID) error {
	return nil
}

func (p *Projects) Search(ctx context.Context, s string) (chan project.Project, error) {
	return nil, nil
}

func (p *Projects) SearchByUsersName(ctx context.Context, names []string) (chan project.Project, error) {
	return nil, nil
}

func (p *Projects) AttachUser(ctx context.Context, o project.Project, uid uuid.UUID) error {
	return nil
}

func (p *Projects) DetachUser(ctx context.Context, o project.Project, uid uuid.UUID) error {
	return nil
}
