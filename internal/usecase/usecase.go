package usecase

import (
	"context"

	"github.com/manattan/mumbai/internal/domain"
	"github.com/manattan/mumbai/internal/domain/model"
)

type UseCase interface {
	CreateUser(ctx context.Context, name, email string) (*model.User, error)
	GetUser(ctx context.Context, id uint) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, limit, offset int) ([]*model.User, error)
}

type userUseCase struct {
	repo domain.Repository
}

var _ UseCase = (*userUseCase)(nil)

func NewUseCase(repo domain.Repository) UseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) CreateUser(ctx context.Context, name, email string) (*model.User, error) {
	user := &model.User{
		Name:  name,
		Email: email,
	}
	return u.repo.User.Create(ctx, user)
}

func (u *userUseCase) GetUser(ctx context.Context, id uint) (*model.User, error) {
	return u.repo.User.GetByID(ctx, id)
}

func (u *userUseCase) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	return u.repo.User.GetByEmail(ctx, email)
}

func (u *userUseCase) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return u.repo.User.Update(ctx, user)
}

func (u *userUseCase) DeleteUser(ctx context.Context, id uint) error {
	return u.repo.User.Delete(ctx, id)
}

func (u *userUseCase) ListUsers(ctx context.Context, limit, offset int) ([]*model.User, error) {
	return u.repo.User.List(ctx, limit, offset)
}
