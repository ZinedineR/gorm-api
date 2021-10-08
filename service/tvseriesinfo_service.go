package service

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilTV occurs when a nil TV is passed.
	ErrNilTV = errors.New("TV is nil")
)

// TVService responsible for any flow related to TV.
// It also implements TVService.
type TVService struct {
	TVRepo TVRepository
}

// NewTVService creates an instance of TVService.
func NewTVService(TVRepo TVRepository) *TVService {
	return &TVService{
		TVRepo: TVRepo,
	}
}

type TVUseCase interface {
	Create(ctx context.Context, TV *entity.TV) error
	GetListTV(ctx context.Context, limit, offset string) ([]*entity.TV, error)
	GetDetailTV(ctx context.Context, ID uuid.UUID) (*entity.TV, error)
	UpdateTV(ctx context.Context, TV *entity.TV) error
	DeleteTV(ctx context.Context, ID uuid.UUID) error
}

type TVRepository interface {
	Insert(ctx context.Context, TV *entity.TV) error
	GetListTV(ctx context.Context, limit, offset string) ([]*entity.TV, error)
	GetDetailTV(ctx context.Context, ID uuid.UUID) (*entity.TV, error)
	UpdateTV(ctx context.Context, TV *entity.TV) error
	DeleteTV(ctx context.Context, ID uuid.UUID) error
}

func (svc TVService) Create(ctx context.Context, TV *entity.TV) error {
	// Checking nil TV
	if TV == nil {
		return ErrNilTV
	}

	// Generate id if nil
	if TV.Id == uuid.Nil {
		TV.Id = uuid.New()
	}

	if err := svc.TVRepo.Insert(ctx, TV); err != nil {
		return errors.Wrap(err, "[TVService-Create]")
	}
	return nil
}

func (svc TVService) GetListTV(ctx context.Context, limit, offset string) ([]*entity.TV, error) {
	TV, err := svc.TVRepo.GetListTV(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[TVService-Create]")
	}
	return TV, nil
}

func (svc TVService) GetDetailTV(ctx context.Context, ID uuid.UUID) (*entity.TV, error) {
	TV, err := svc.TVRepo.GetDetailTV(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[TVService-Create]")
	}
	return TV, nil
}

func (svc TVService) DeleteTV(ctx context.Context, ID uuid.UUID) error {
	err := svc.TVRepo.DeleteTV(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[TVService-Create]")
	}
	return nil
}

func (svc TVService) UpdateTV(ctx context.Context, TV *entity.TV) error {
	// Checking nil TV
	if TV == nil {
		return ErrNilTV
	}

	// Generate id if nil
	if TV.Id == uuid.Nil {
		TV.Id = uuid.New()
	}

	if err := svc.TVRepo.UpdateTV(ctx, TV); err != nil {
		return errors.Wrap(err, "[TVService-Create]")
	}
	return nil
}
