package repository

import (
	"github.com/manattan/mumbai/internal/domain"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) domain.Repository {
	return domain.Repository{
		User: newUserRepository(db),
	}
}
