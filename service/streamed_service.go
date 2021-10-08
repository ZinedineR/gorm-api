package service

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilStreamed occurs when a nil Streamed is passed.
	ErrNilStreamed = errors.New("Streamed is nil")
)

// StreamedService responsible for any flow related to Streamed.
// It also implements StreamedService.
type StreamedService struct {
	StreamedRepo StreamedRepository
}

// NewStreamedService creates an instance of StreamedService.
func NewStreamedService(StreamedRepo StreamedRepository) *StreamedService {
	return &StreamedService{
		StreamedRepo: StreamedRepo,
	}
}

type StreamedUseCase interface {
	Create(ctx context.Context, Streamed *entity.Streamed) error
	GetListStreamed(ctx context.Context, limit, offset string) ([]*entity.Streamed, error)
	GetDetailStreamed(ctx context.Context, ID uuid.UUID) (*entity.Streamed, error)
	UpdateStreamed(ctx context.Context, Streamed *entity.Streamed) error
	DeleteStreamed(ctx context.Context, ID uuid.UUID) error
}

type StreamedRepository interface {
	Insert(ctx context.Context, Streamed *entity.Streamed) error
	GetListStreamed(ctx context.Context, limit, offset string) ([]*entity.Streamed, error)
	GetDetailStreamed(ctx context.Context, ID uuid.UUID) (*entity.Streamed, error)
	UpdateStreamed(ctx context.Context, Streamed *entity.Streamed) error
	DeleteStreamed(ctx context.Context, ID uuid.UUID) error
}

func (svc StreamedService) Create(ctx context.Context, Streamed *entity.Streamed) error {
	// Checking nil Streamed
	if Streamed == nil {
		return ErrNilStreamed
	}

	// Generate id if nil
	if Streamed.Id == uuid.Nil {
		Streamed.Id = uuid.New()
	}

	if err := svc.StreamedRepo.Insert(ctx, Streamed); err != nil {
		return errors.Wrap(err, "[StreamedService-Create]")
	}
	return nil
}

func (svc StreamedService) GetListStreamed(ctx context.Context, limit, offset string) ([]*entity.Streamed, error) {
	Streamed, err := svc.StreamedRepo.GetListStreamed(ctx, limit, offset)
	if err != nil {
		return nil, errors.Wrap(err, "[StreamedService-Create]")
	}
	return Streamed, nil
}

func (svc StreamedService) GetDetailStreamed(ctx context.Context, ID uuid.UUID) (*entity.Streamed, error) {
	Streamed, err := svc.StreamedRepo.GetDetailStreamed(ctx, ID)
	if err != nil {
		return nil, errors.Wrap(err, "[StreamedService-Create]")
	}
	return Streamed, nil
}

func (svc StreamedService) DeleteStreamed(ctx context.Context, ID uuid.UUID) error {
	err := svc.StreamedRepo.DeleteStreamed(ctx, ID)
	if err != nil {
		return errors.Wrap(err, "[StreamedService-Create]")
	}
	return nil
}

func (svc StreamedService) UpdateStreamed(ctx context.Context, Streamed *entity.Streamed) error {
	// Checking nil Streamed
	if Streamed == nil {
		return ErrNilStreamed
	}

	// Generate id if nil
	if Streamed.Id == uuid.Nil {
		Streamed.Id = uuid.New()
	}

	if err := svc.StreamedRepo.UpdateStreamed(ctx, Streamed); err != nil {
		return errors.Wrap(err, "[StreamedService-Create]")
	}
	return nil
}
