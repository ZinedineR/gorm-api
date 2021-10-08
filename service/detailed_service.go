package service

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilDetailed occurs when a nil Detailed is passed.
	ErrNilDetailed = errors.New("Detailed is nil")
)

// DetailedService responsible for any flow related to Detailed.
// It also implements DetailedService.
type DetailedService struct {
	DetailedRepo DetailedRepository
}

// NewDetailedService creates an instance of DetailedService.
func NewDetailedService(DetailedRepo DetailedRepository) *DetailedService {
	return &DetailedService{
		DetailedRepo: DetailedRepo,
	}
}

type DetailedUseCase interface {
	Create(ctx context.Context, Detailed *entity.Detailed) error
	GetListDetailed(ctx context.Context, limit, offset string) ([]*entity.Detailed, error)
	GetDetailDetailed(ctx context.Context, ID uuid.UUID) (*entity.Detailed, error)
	UpdateDetailed(ctx context.Context, Detailed *entity.Detailed) error
	DeleteDetailed(ctx context.Context, ID uuid.UUID) error
}

type DetailedRepository interface {
	Insert(ctx context.Context, Detailed *entity.Detailed) error
	GetListDetailed(ctx context.Context, limit, offset string) ([]*entity.Detailed, error)
	GetDetailDetailed(ctx context.Context, ID uuid.UUID) (*entity.Detailed, error)
	UpdateDetailed(ctx context.Context, Detailed *entity.Detailed) error
	DeleteDetailed(ctx context.Context, ID uuid.UUID) error
}

func (svc DetailedService) Create(ctx context.Context, Detailed *entity.Detailed) error {
	// Checking nil Detailed
	if Detailed == nil {
		return ErrNilDetailed
	}

	// Generate id if nil
	if Detailed.Id == uuid.Nil {
		Detailed.Id = uuid.New()
	}

	if err := svc.DetailedRepo.Insert(ctx, Detailed); err != nil {
		return errors.Wrap(err, "[DetailedService-Create]")
	}
	return nil
}

func (svc DetailedService) GetListDetailed(ctx context.Context, limit, offset string) ([]*entity.Detailed, error) {
	Detailed, err := svc.DetailedRepo.GetListDetailed(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[DetailedService-Create]")
	}
	return Detailed, nil
}

func (svc DetailedService) GetDetailDetailed(ctx context.Context, ID uuid.UUID) (*entity.Detailed, error) {
	Detailed, err := svc.DetailedRepo.GetDetailDetailed(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[DetailedService-Create]")
	}
	return Detailed, nil
}

func (svc DetailedService) DeleteDetailed(ctx context.Context, ID uuid.UUID) error {
	err := svc.DetailedRepo.DeleteDetailed(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[DetailedService-Create]")
	}
	return nil
}

func (svc DetailedService) UpdateDetailed(ctx context.Context, Detailed *entity.Detailed) error {
	// Checking nil Detailed
	if Detailed == nil {
		return ErrNilDetailed
	}

	// Generate id if nil
	if Detailed.Id == uuid.Nil {
		Detailed.Id = uuid.New()
	}

	if err := svc.DetailedRepo.UpdateDetailed(ctx, Detailed); err != nil {
		return errors.Wrap(err, "[DetailedService-Create]")
	}
	return nil
}
