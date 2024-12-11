package gorm

import (
	"github.com/jkitajima/efm/identity/auth/pkg/user"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewRepo(db *gorm.DB) user.Repoer {
	return &DB{db}
}
