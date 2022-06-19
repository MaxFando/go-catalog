package organizationrepo

import (
	"context"
	"github.com/MaxFando/go-catalog/internal/entities/organization"
	"github.com/google/uuid"
)

type OrganizationStore interface {
	Create(ctx context.Context, org organization.Organization) (*uuid.UUID, error)
	Read(ctx context.Context, orgId uuid.UUID) (*uuid.UUID, error)
	Delete(ctx context.Context, orgId uuid.UUID) error
	Search(ctx context.Context, s string) (chan organization.Organization, error)
	SearchByUsersName(ctx context.Context, names []string) (chan organization.Organization, error)

	AttachUser(ctx context.Context, org organization.Organization, uid uuid.UUID) error
	DetachUser(ctx context.Context, org organization.Organization, uid uuid.UUID) error
}

type Organizations struct {
	oStore OrganizationStore
}

func New(oStore OrganizationStore) *Organizations {
	return &Organizations{
		oStore: oStore,
	}
}

func (o *Organizations) Create(ctx context.Context, org organization.Organization) (*uuid.UUID, error) {
	return nil, nil
}

func (o *Organizations) Read(ctx context.Context, orgId uuid.UUID) (*uuid.UUID, error) {
	return nil, nil
}

func (o *Organizations) Delete(ctx context.Context, orgId uuid.UUID) error {
	return nil
}

func (o *Organizations) Search(ctx context.Context, s string) (chan organization.Organization, error) {
	return nil, nil
}

func (o *Organizations) SearchByUsersName(ctx context.Context, names []string) (chan organization.Organization, error) {
	return nil, nil
}

func (o *Organizations) AttachUser(ctx context.Context, org organization.Organization, uid uuid.UUID) error {
	return nil
}

func (o *Organizations) DetachUser(ctx context.Context, org organization.Organization, uid uuid.UUID) error {
	return nil
}
