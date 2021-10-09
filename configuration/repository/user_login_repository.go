package repository

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// UserRepository connects entity.User with database.
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository creates an instance of RoleRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Insert inserts User data to database.
func (repo *UserRepository) Insert(ctx context.Context, ent *entity.User) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[UserRepository-Insert]")
	}
	return nil
}
