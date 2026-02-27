package service

import (
	"context"
	"strings"

	"github.com/billalhossainjoy/openparadox/internal/domain"
	"github.com/billalhossainjoy/openparadox/internal/repository"
	"github.com/google/uuid"
)

type UserService struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func(s *UserService) CreateUser(ctx context.Context, name, email string) (domain.User, error) {
	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	if name=="" || email==""|| !strings.Contains(email, "@") {
		return  domain.User{}, domain.ErrInvalidInput
	}
	
	user:= domain.User{
		ID: uuid.NewString(),
		Name: name,
		Email: email,
	}

	if err := s.repo.Save(ctx, user); err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *UserService) GetUsers(ctx context.Context) ([]domain.User, error) {
	return s.repo.FindAll(ctx)
}

func (s *UserService) GetUser(ctx context.Context, id string) (domain.User, error) {
	return s.repo.FindById(ctx, id)
}
