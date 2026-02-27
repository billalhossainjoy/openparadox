package repository

import (
	"context"
	"sync"

	"github.com/billalhossainjoy/openparadox/internal/domain"
)


type UserRepository interface {
	Save(ctx context.Context, user domain.User) error
	FindAll(ctx context.Context) ([]domain.User, error)
	FindById(ctx context.Context, id string) (domain.User, error)
}

type InMemoryUserRepository struct{
	data map[string]domain.User
	mu sync.RWMutex
}

func NewMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		data: make(map[string]domain.User),
	}
}

func (r *InMemoryUserRepository) Save(ctx context.Context, user domain.User) error{
	_ = ctx
	r.mu.Lock()
	defer r.mu.Unlock()

	r.data[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) FindAll(ctx context.Context, ) ([]domain.User, error) {
	_ = ctx
	r.mu.Lock()
	defer r.mu.Unlock()

	users := []domain.User{}
	for _, u:= range r.data {
		users = append(users, u)
	}
	return users, nil
}

func (r *InMemoryUserRepository) FindById (ctx context.Context, id string) (domain.User, error) {
	_ = ctx
	r.mu.Lock()
	defer r.mu.Unlock()

	user, ok := r.data[id]
	if !ok{
		return domain.User{}, domain.ErrNotFound
	}

	return user, nil	
}

