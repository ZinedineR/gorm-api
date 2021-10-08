package repository

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// DetailedRepository connects entity.Detailed with database.
type DetailedRepository struct {
	db *gorm.DB
}

// NewDetailedRepository creates an instance of RoleRepository.
func NewDetailedRepository(db *gorm.DB) *DetailedRepository {
	return &DetailedRepository{
		db: db,
	}
}

// Insert inserts Detailed data to database.
func (repo *DetailedRepository) Insert(ctx context.Context, ent *entity.Detailed) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Detailed{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[DetailedRepository-Insert]")
	}
	return nil
}

func (repo *DetailedRepository) GetListDetailed(ctx context.Context, limit, offset string) ([]*entity.Detailed, error) {
	var models []*entity.Detailed
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Detailed{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[DetailedRepository-FindAll]")
	}
	return models, nil
}

func (repo *DetailedRepository) GetDetailDetailed(ctx context.Context, ID uuid.UUID) (*entity.Detailed, error) {
	var models *entity.Detailed
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Detailed{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[DetailedRepository-FindById]")
	}
	return models, nil
}

func (repo *DetailedRepository) DeleteDetailed(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Detailed{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[DetailedRepository-Delete]")
	}
	return nil
}

func (repo *DetailedRepository) UpdateDetailed(ctx context.Context, ent *entity.Detailed) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Detailed{Id: ent.Id}).
		Select("title", "producer").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[DetailedRepository-Update]")
	}
	return nil
}
