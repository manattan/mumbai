package repository

import (
	"context"
	"time"

	"github.com/manattan/mumbai/internal/domain"
	"github.com/manattan/mumbai/internal/domain/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

var _ domain.UserRepository = (*userRepository)(nil)

func newUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

type UserModel struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserModel) TableName() string {
	return "users"
}

func (u *UserModel) ToEntity() *model.User {
	return &model.User{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func FromEntity(user *model.User) *UserModel {
	return &UserModel{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	userModel := FromEntity(user)
	if err := r.db.WithContext(ctx).Create(userModel).Error; err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var userModel UserModel
	if err := r.db.WithContext(ctx).First(&userModel, id).Error; err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var userModel UserModel
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&userModel).Error; err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	userModel := FromEntity(user)
	if err := r.db.WithContext(ctx).Save(userModel).Error; err != nil {
		return nil, err
	}
	return userModel.ToEntity(), nil
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&UserModel{}, id).Error
}

func (r *userRepository) List(ctx context.Context, limit, offset int) ([]*model.User, error) {
	var userModels []UserModel
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&userModels).Error; err != nil {
		return nil, err
	}

	users := make([]*model.User, len(userModels))
	for i, userModel := range userModels {
		users[i] = userModel.ToEntity()
	}
	return users, nil
}
