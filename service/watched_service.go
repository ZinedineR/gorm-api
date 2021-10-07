package service

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilWatched occurs when a nil Watched is passed.
	ErrNilWatched = errors.New("Watched is nil")
)

// WatchedService responsible for any flow related to Watched.
// It also implements WatchedService.
type WatchedService struct {
	WatchedRepo WatchedRepository
}

// NewWatchedService creates an instance of WatchedService.
func NewWatchedService(WatchedRepo WatchedRepository) *WatchedService {
	return &WatchedService{
		WatchedRepo: WatchedRepo,
	}
}

type WatchedUseCase interface {
	Create(ctx context.Context, Watched *entity.Watched) error
	GetListWatched(ctx context.Context, limit, offset string) ([]*entity.Watched, error)
	GetDetailWatched(ctx context.Context, ID uuid.UUID) (*entity.Watched, error)
	UpdateWatched(ctx context.Context, Watched *entity.Watched) error
	DeleteWatched(ctx context.Context, ID uuid.UUID) error
}

type WatchedRepository interface {
	Insert(ctx context.Context, Watched *entity.Watched) error
	GetListWatched(ctx context.Context, limit, offset string) ([]*entity.Watched, error)
	GetDetailWatched(ctx context.Context, ID uuid.UUID) (*entity.Watched, error)
	UpdateWatched(ctx context.Context, Watched *entity.Watched) error
	DeleteWatched(ctx context.Context, ID uuid.UUID) error
}

func (svc WatchedService) Create(ctx context.Context, Watched *entity.Watched) error {
	// Checking nil Watched
	if Watched == nil {
		return ErrNilWatched
	}

	// Generate id if nil
	if Watched.Id == uuid.Nil {
		Watched.Id = uuid.New()
	}

	if err := svc.WatchedRepo.Insert(ctx, Watched); err != nil {
		return errors.Wrap(err, "[WatchedService-Create]")
	}
	return nil
}

func (svc WatchedService) GetListWatched(ctx context.Context, limit, offset string) ([]*entity.Watched, error) {
	Watched, err := svc.WatchedRepo.GetListWatched(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[WatchedService-Create]")
	}
	return Watched, nil
}

func (svc WatchedService) GetDetailWatched(ctx context.Context, ID uuid.UUID) (*entity.Watched, error) {
	Watched, err := svc.WatchedRepo.GetDetailWatched(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[WatchedService-Create]")
	}
	return Watched, nil
}

func (svc WatchedService) DeleteWatched(ctx context.Context, ID uuid.UUID) error {
	err := svc.WatchedRepo.DeleteWatched(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[WatchedService-Create]")
	}
	return nil
}

func (svc WatchedService) UpdateWatched(ctx context.Context, Watched *entity.Watched) error {
	// Checking nil Watched
	if Watched == nil {
		return ErrNilWatched
	}

	// Generate id if nil
	if Watched.Id == uuid.Nil {
		Watched.Id = uuid.New()
	}

	if err := svc.WatchedRepo.UpdateWatched(ctx, Watched); err != nil {
		return errors.Wrap(err, "[WatchedService-Create]")
	}
	return nil
}
