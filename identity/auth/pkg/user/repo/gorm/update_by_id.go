package gorm

import (
	"context"

	"github.com/google/uuid"
	"github.com/jkitajima/efm/identity/auth/pkg/user"
)

func (db *DB) UpdateByID(ctx context.Context, id uuid.UUID, user *user.User) error {

	result := db.Model(&UserModel{}).Where("id = ?", id.String()).Update("email_verified", true)
	// db.First(user, "id = ?", id.String())

	return result.Error
}
