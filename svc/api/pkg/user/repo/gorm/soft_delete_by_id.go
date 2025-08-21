package gorm

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jkitajima/efm/svc/api/pkg/user"
)

func (db *DB) SoftDeleteByID(ctx context.Context, id uuid.UUID) error {
	model := UserModel{ID: id}
	result := db.Delete(&model)
	if result.Error != nil {
		fmt.Println(result.Error)
		return user.ErrInternal
	}

	return nil
}
