package organization

import (
	"github.com/MaxFando/go-catalog/internal/entities/user"
	"github.com/google/uuid"
)

type Organization struct {
	ID    uuid.UUID
	Users []user.User
}
