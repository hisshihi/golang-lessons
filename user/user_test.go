package user

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestService_CreateUser(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		user    User
		wantErr bool
		setup   func(up *MockUserProvider, uc *MockUserCreator, n *MockEventNotifier, user User)
	}{
		{
			name:    "base test",
			user:    User{Email: "hissoffc@gmail.com"},
			wantErr: false,
			setup: func(up *MockUserProvider, uc *MockUserCreator, n *MockEventNotifier, user User) {
				up.
					On("User", mock.Anything, user.Email).
					Once().
					Return(nil, nil)
				uc.
					On("Create", mock.Anything, user).
					Once().
					Return(1, nil)
				n.
					On("NotifyUserCreated", mock.Anything, user).
					Once().
					Return(nil)
			},
		},
		{
			name:    "user already exists",
			user:    User{Email: "hissoffc@gmail.com"},
			wantErr: true,
			setup: func(up *MockUserProvider, uc *MockUserCreator, n *MockEventNotifier, user User) {
				up.
					On("User", mock.Anything, user.Email).
					Once().
					Return(&user, nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			up := NewMockUserProvider(t)
			uc := NewMockUserCreator(t)
			n := NewMockEventNotifier(t)

			tt.setup(up, uc, n, tt.user)

			s := &Service{up, uc, n}
			_, gotErr := s.CreateUser(context.Background(), tt.user)
			if gotErr != nil {
				if !tt.wantErr {
					t.Fatalf("CreateUser() filed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateUser() succeeded unexpectedly")
			}
		})
	}
}
