package gorm

import (
	"context"

	"github.com/jkitajima/efm/svc/api/pkg/user"
	"gorm.io/gorm"
)

func (db *DB) Insert(ctx context.Context, u *user.User) error {
	model := &UserModel{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Role:      u.Role,
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
