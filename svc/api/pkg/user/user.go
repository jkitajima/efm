package user

import (
	"context"
	"database/sql/driver"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	Default Role = "default"
	Admin   Role = "admin"
)

func (r *Role) Scan(src any) error {
	*r = Role(src.([]byte))
	return nil
}

func (r *Role) Value() (driver.Value, error) {
	return string(*r), nil
}

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  *string
	Role      Role
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Service struct {
	Repo Repoer
}

type Repoer interface {
	Insert(context.Context, *User) error
	// FindByID(context.Context, uuid.UUID) (*User, error)
	// UpdateByID(context.Context, uuid.UUID, *User) error
	// DeleteByID(context.Context, uuid.UUID) error
}

var (
	ErrInternal = errors.New("api: the user service encountered an unexpected condition that prevented it from fulfilling the request")
)
