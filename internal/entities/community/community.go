package community

import (
	"github.com/MaxFando/go-catalog/internal/entities/user"
	"github.com/google/uuid"
)

type Community struct {
	ID    uuid.UUID
	Users []user.User
}
