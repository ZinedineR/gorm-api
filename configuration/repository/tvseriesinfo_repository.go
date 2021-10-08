package repository

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// TVRepository connects entity.TV with database.
type TVRepository struct {
	db *gorm.DB
}

// NewTVRepository creates an instance of RoleRepository.
func NewTVRepository(db *gorm.DB) *TVRepository {
	return &TVRepository{
		db: db,
	}
}

// Insert inserts TV data to database.
func (repo *TVRepository) Insert(ctx context.Context, ent *entity.TV) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.TV{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[TVRepository-Insert]")
	}
	return nil
}

func (repo *TVRepository) GetListTV(ctx context.Context, limit, offset string) ([]*entity.TV, error) {
	var models []*entity.TV
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.TV{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[TVRepository-FindAll]")
	}
	return models, nil
}

func (repo *TVRepository) GetDetailTV(ctx context.Context, ID uuid.UUID) (*entity.TV, error) {
	var models *entity.TV
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.TV{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[TVRepository-FindById]")
	}
	return models, nil
}

func (repo *TVRepository) DeleteTV(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.TV{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[TVRepository-Delete]")
	}
	return nil
}

func (repo *TVRepository) UpdateTV(ctx context.Context, ent *entity.TV) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.TV{Id: ent.Id}).
		Select("title", "producer").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[TVRepository-Update]")
	}
	return nil
}
