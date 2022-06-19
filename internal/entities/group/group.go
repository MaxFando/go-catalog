package group

import (
	"github.com/MaxFando/go-catalog/internal/entities/user"
	"github.com/google/uuid"
)

type Group struct {
	ID    uuid.UUID
	Users []user.User
}
