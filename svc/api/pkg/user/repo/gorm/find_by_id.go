package gorm

// import (
// 	"context"

// 	"github.com/google/uuid"
// )

// func (db *DB) FindByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
// 	var model UserModel
// 	result := db.First(&model, "id = ?", id.String())

// 	// user := user.User{
// 		// ID: model.ID,
// 		// Email: model.Email,
// 		// EmailVerified:              model.EmailVerified,
// 		// Password: model.Password,
// 		// VerificationCode:           model.VerificationCode,
// 		// VerificationCodeExpiration: model.VerificationCodeExpiration,
// 		// CreatedAt: model.CreatedAt,
// 		// UpdatedAt: model.UpdatedAt,
// 		// DeletedAt: &model.DeletedAt.Time,
// 	// }

// 	return &user, result.Error
// }
