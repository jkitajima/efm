package gorm

import (
	"context"

	"github.com/jkitajima/efm/identity/auth/pkg/user"
	"gorm.io/gorm"
)

func (db *DB) Insert(ctx context.Context, u *user.User) error {
	model := &UserModel{
		Email: u.Email,
		// EmailVerified:              u.EmailVerified,
		Password: u.Password,
		// VerificationCode:           u.VerificationCode,
		// VerificationCodeExpiration: u.VerificationCodeExpiration,
	}

	if u.DeletedAt != nil {
		model.DeletedAt = gorm.DeletedAt{
			Time:  *u.DeletedAt,
			Valid: true,
		}
	}

	result := db.Create(model)

	u.ID = model.ID
	u.CreatedAt = model.CreatedAt
	u.UpdatedAt = model.UpdatedAt

	return result.Error
}
