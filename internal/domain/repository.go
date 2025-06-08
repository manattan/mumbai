package domain

import (
	"context"

	"github.com/manattan/mumbai/internal/domain/model"
)

type Repository struct {
	User UserRepository
}

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetByID(ctx context.Context, id uint) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, limit, offset int) ([]*model.User, error)
}
