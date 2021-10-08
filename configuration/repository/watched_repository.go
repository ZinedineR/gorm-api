package repository

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// WatchedRepository connects entity.Watched with database.
type WatchedRepository struct {
	db *gorm.DB
}

// NewWatchedRepository creates an instance of RoleRepository.
func NewWatchedRepository(db *gorm.DB) *WatchedRepository {
	return &WatchedRepository{
		db: db,
	}
}

// Insert inserts Watched data to database.
func (repo *WatchedRepository) Insert(ctx context.Context, ent *entity.Watched) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Watched{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[WatchedRepository-Insert]")
	}
	return nil
}

func (repo *WatchedRepository) GetListWatched(ctx context.Context, limit, offset string) ([]*entity.Watched, error) {
	var models []*entity.Watched
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Watched{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[WatchedRepository-FindAll]")
	}
	return models, nil
}

func (repo *WatchedRepository) GetDetailWatched(ctx context.Context, ID uuid.UUID) (*entity.Watched, error) {
	var models *entity.Watched
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Watched{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[WatchedRepository-FindById]")
	}
	return models, nil
}

func (repo *WatchedRepository) DeleteWatched(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Watched{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[WatchedRepository-Delete]")
	}
	return nil
}

func (repo *WatchedRepository) UpdateWatched(ctx context.Context, ent *entity.Watched) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Watched{Id: ent.Id}).
		Select("Watched_id", "season", "episodes").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[WatchedRepository-Update]")
	}
	return nil
}
