package user

import (
	"context"

	"github.com/hussar_company/vstazherstve_back/entity"
)

// Service инкапсулирует логику юзкейсов для пользователей
type Service interface {
	Get(ctx context.Context, id string) (User, error)

	Create(ctx context.Context, input CreateUserRequest) error
}

// Пользователь
type User struct {
	entity.User
}

// Создание пользователя
type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type service struct {
	repo Repository
}

// Создает новый сервис пользователей
func NewService(repo Repository) Service {
	return service{repo}
}

// Возвращает пользователя с указанным идентификатором
func (s service) Get(ctx context.Context, id string) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return User{user}, nil
}

// Создает нового пользователя
func (s service) Create(ctx context.Context, req CreateUserRequest) error {
	err := s.repo.Create(ctx, entity.User{
		Name:     req.Name,
		Password: req.Password,
	})
	return err
}
