package organizationememstore

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/organization"
	"github.com/MaxFando/go-catalog/internal/logic/app/repos/organizationrepo"
	"github.com/google/uuid"
	"sync"
)

var _ organizationrepo.OrganizationStore = &Organizations{}

type Organizations struct {
	sync.Mutex
	m map[uuid.UUID]organization.Organization
}

func New() *Organizations {
	return &Organizations{
		m: make(map[uuid.UUID]organization.Organization),
	}
}

func (org *Organizations) Create(ctx context.Context, o organization.Organization) (*uuid.UUID, error) {
	return nil, nil
}

func (org *Organizations) Read(ctx context.Context, oId uuid.UUID) (*uuid.UUID, error) {
	return nil, nil
}

func (org *Organizations) Delete(ctx context.Context, oId uuid.UUID) error {
	return nil
}

func (org *Organizations) Search(ctx context.Context, s string) (chan organization.Organization, error) {
	return nil, nil
}

func (org *Organizations) SearchByUsersName(ctx context.Context, names []string) (chan organization.Organization, error) {
	return nil, nil
}

func (org *Organizations) AttachUser(ctx context.Context, o organization.Organization, uid uuid.UUID) error {
	return nil
}

func (org *Organizations) DetachUser(ctx context.Context, o organization.Organization, uid uuid.UUID) error {
	return nil
}
