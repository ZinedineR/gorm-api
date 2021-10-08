package service

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilActor occurs when a nil Actor is passed.
	ErrNilActor = errors.New("Actor is nil")
)

// ActorService responsible for any flow related to Actor.
// It also implements ActorService.
type ActorService struct {
	ActorRepo ActorRepository
}

// NewActorService creates an instance of ActorService.
func NewActorService(ActorRepo ActorRepository) *ActorService {
	return &ActorService{
		ActorRepo: ActorRepo,
	}
}

type ActorUseCase interface {
	Create(ctx context.Context, Actor *entity.Actor) error
	GetListActor(ctx context.Context, limit, offset string) ([]*entity.Actor, error)
	GetDetailActor(ctx context.Context, ID uuid.UUID) (*entity.Actor, error)
	UpdateActor(ctx context.Context, Actor *entity.Actor) error
	DeleteActor(ctx context.Context, ID uuid.UUID) error
}

type ActorRepository interface {
	Insert(ctx context.Context, Actor *entity.Actor) error
	GetListActor(ctx context.Context, limit, offset string) ([]*entity.Actor, error)
	GetDetailActor(ctx context.Context, ID uuid.UUID) (*entity.Actor, error)
	UpdateActor(ctx context.Context, Actor *entity.Actor) error
	DeleteActor(ctx context.Context, ID uuid.UUID) error
}

func (svc ActorService) Create(ctx context.Context, Actor *entity.Actor) error {
	// Checking nil Actor
	if Actor == nil {
		return ErrNilActor
	}

	// Generate id if nil
	if Actor.Id == uuid.Nil {
		Actor.Id = uuid.New()
	}

	if err := svc.ActorRepo.Insert(ctx, Actor); err != nil {
		return errors.Wrap(err, "[ActorService-Create]")
	}
	return nil
}

func (svc ActorService) GetListActor(ctx context.Context, limit, offset string) ([]*entity.Actor, error) {
	Actor, err := svc.ActorRepo.GetListActor(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[ActorService-Create]")
	}
	return Actor, nil
}

func (svc ActorService) GetDetailActor(ctx context.Context, ID uuid.UUID) (*entity.Actor, error) {
	Actor, err := svc.ActorRepo.GetDetailActor(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[ActorService-Create]")
	}
	return Actor, nil
}

func (svc ActorService) DeleteActor(ctx context.Context, ID uuid.UUID) error {
	err := svc.ActorRepo.DeleteActor(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[ActorService-Create]")
	}
	return nil
}

func (svc ActorService) UpdateActor(ctx context.Context, Actor *entity.Actor) error {
	// Checking nil Actor
	if Actor == nil {
		return ErrNilActor
	}

	// Generate id if nil
	if Actor.Id == uuid.Nil {
		Actor.Id = uuid.New()
	}

	if err := svc.ActorRepo.UpdateActor(ctx, Actor); err != nil {
		return errors.Wrap(err, "[ActorService-Create]")
	}
	return nil
}
