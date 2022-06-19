package project

import (
	"github.com/MaxFando/go-catalog/internal/entities/user"
	"github.com/google/uuid"
)

type Project struct {
	ID    uuid.UUID
	Users []user.User
}
