package user

import (
	"github.com/MaxFando/go-catalog/internal/entities/community"
	"github.com/MaxFando/go-catalog/internal/entities/group"
	"github.com/MaxFando/go-catalog/internal/entities/organization"
	"github.com/MaxFando/go-catalog/internal/entities/project"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Name        string
	Data        string
	Permissions int

	Group        *group.Group
	Organization *organization.Organization
	Project      *project.Project
	Society      *community.Community
}
