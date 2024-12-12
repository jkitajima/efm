package user

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID    uuid.UUID
	Email string
	// EmailVerified              bool
	Password string
	// VerificationCode           *string
	// VerificationCodeExpiration *int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Service struct {
	Repo Repoer
}

type Repoer interface {
	Insert(context.Context, *User) error
	FindByID(context.Context, uuid.UUID) (*User, error)
	UpdateByID(context.Context, uuid.UUID, *User) error
	// DeleteByID(context.Context, uuid.UUID) error
}

var (
	ErrInternal = errors.New("auth: the user service encountered an unexpected condition that prevented it from fulfilling the request")
)
