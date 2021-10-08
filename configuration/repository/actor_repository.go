package repository

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// ActorRepository connects entity.Actor with database.
type ActorRepository struct {
	db *gorm.DB
}

// NewActorRepository creates an instance of RoleRepository.
func NewActorRepository(db *gorm.DB) *ActorRepository {
	return &ActorRepository{
		db: db,
	}
}

// Insert inserts Actor data to database.
func (repo *ActorRepository) Insert(ctx context.Context, ent *entity.Actor) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Actor{}).
		Create(ent).
		Error; err != nil {
		return errors.Wrap(err, "[ActorRepository-Insert]")
	}
	return nil
}

func (repo *ActorRepository) GetListActor(ctx context.Context, limit, offset string) ([]*entity.Actor, error) {
	var models []*entity.Actor
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Actor{}).
		Find(&models).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ActorRepository-FindAll]")
	}
	return models, nil
}

func (repo *ActorRepository) GetDetailActor(ctx context.Context, ID uuid.UUID) (*entity.Actor, error) {
	var models *entity.Actor
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Actor{}).
		Take(&models, ID).
		Error; err != nil {
		return nil, errors.Wrap(err, "[ActorRepository-FindById]")
	}
	return models, nil
}

func (repo *ActorRepository) DeleteActor(ctx context.Context, ID uuid.UUID) error {
	if err := repo.db.
		WithContext(ctx).
		Delete(&entity.Actor{Id: ID}).Error; err != nil {
		return errors.Wrap(err, "[ActorRepository-Delete]")
	}
	return nil
}

func (repo *ActorRepository) UpdateActor(ctx context.Context, ent *entity.Actor) error {
	if err := repo.db.
		WithContext(ctx).
		Model(&entity.Actor{Id: ent.Id}).
		Select("Actor_id", "season", "episodes").
		Updates(ent).Error; err != nil {
		return errors.Wrap(err, "[ActorRepository-Update]")
	}
	return nil
}
