package service

import (
	"api-gorm-setting/entity"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	// ErrNilUser occurs when a nil User is passed.
	ErrNilUser = errors.New("User is nil")
)

// UserService responsible for any flow related to User.
// It also implements UserService.
type UserService struct {
	UserRepo UserRepository
}

// NewUserService creates an instance of UserService.
func NewUserService(UserRepo UserRepository) *UserService {
	return &UserService{
		UserRepo: UserRepo,
	}
}

type UserUseCase interface {
	Create(ctx context.Context, User *entity.User) error
}

type UserRepository interface {
	Insert(ctx context.Context, User *entity.User) error
}

func (svc UserService) Create(ctx context.Context, User *entity.User) error {
	// Checking nil User
	if User == nil {
		return ErrNilUser
	}

	// Generate id if nil
	if User.Id == uuid.Nil {
		User.Id = uuid.New()
	}

	if err := svc.UserRepo.Insert(ctx, User); err != nil {
		return errors.Wrap(err, "[UserService-Create]")
	}
	return nil
}
