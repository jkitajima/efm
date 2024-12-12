package gorm

import (
	"github.com/jkitajima/efm/svc/api/pkg/user"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func NewRepo(db *gorm.DB) user.Repoer {
	return &DB{db}
}
