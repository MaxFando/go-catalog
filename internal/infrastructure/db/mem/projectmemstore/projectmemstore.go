package projectmemstore

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/project"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/projectrepo"
	"github.com/google/uuid"
	"sync"
)

var _ projectrepo.ProjectStore = &Projects{}

type Projects struct {
	sync.Mutex
	m map[uuid.UUID]project.Project
}

func New() *Projects {
	return &Projects{
		m: make(map[uuid.UUID]project.Project),
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
