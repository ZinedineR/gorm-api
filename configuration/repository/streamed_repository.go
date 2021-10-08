package repository

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// StreamedRepository connects entity.Streamed with database.
type StreamedRepository struct {
	db *gorm.DB
}

// NewStreamedRepository creates an instance of RoleRepository.
func NewStreamedRepository(db *gorm.DB) *StreamedRepository {
	return &StreamedRepository{
		db: db,
	}
}

// Insert inserts Streamed data to database.
func (repo *StreamedRepository) Insert(ctx context.Context, ent *entity.Streamed) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Streamed{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[StreamedRepository-Insert]")
	}
	return nil
}

func (repo *StreamedRepository) GetListStreamed(ctx context.Context, limit, offset string) ([]*entity.Streamed, error) {
	var models []*entity.Streamed
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Streamed{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[StreamedRepository-FindAll]")
	}
	return models, nil
}

func (repo *StreamedRepository) GetDetailStreamed(ctx context.Context, ID uuid.UUID) (*entity.Streamed, error) {
	var models *entity.Streamed
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Streamed{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[StreamedRepository-FindById]")
	}
	return models, nil
}

func (repo *StreamedRepository) DeleteStreamed(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Streamed{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[StreamedRepository-Delete]")
	}
	return nil
}

func (repo *StreamedRepository) UpdateStreamed(ctx context.Context, ent *entity.Streamed) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Streamed{Id: ent.Id}).
		Select("streamed_id", "platform").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[StreamedRepository-Update]")
	}
	return nil
}
