package user

import (
	"context"
	"errors"
	"fmt"
)

type User struct {
	Email string
}

//go:generate go run github.com/vektra/mockery/v3@v3.7.0
type UserCreator interface {
	Create(ctx context.Context, u User) (int, error)
}

type UserProvider interface {
	User(ctx context.Context, email string) (*User, error)
}

type EventNotifier interface {
	NotifyUserCreated(ctx context.Context, user User) error
}

type Service struct {
	userProvider  UserProvider
	userCreator   UserCreator
	eventNotifier EventNotifier
}

func (s *Service) CreateUser(ctx context.Context, user User) (int, error) {
	foundUser, err := s.userProvider.User(ctx, user.Email)
	if err != nil {
		return 0, fmt.Errorf("can`t get user: %w", err)
	}

	if foundUser != nil {
		return 0, errors.New("user already exists")
	}

	uid, err := s.userCreator.Create(ctx, user)
	if err != nil {
		return 0, fmt.Errorf("can`t create user: %w", err)
	}

	if err := s.eventNotifier.NotifyUserCreated(ctx, user); err != nil {
		return 0, fmt.Errorf("can`t notify user created: %w", err)
	}

	return uid, err
}
