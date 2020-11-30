package account

import (
	"context"
)

type Token struct {
	AccessToken string
	RefreshToke string
	ExpiresIn   int64
}

type SessionManager interface {
	New(ctx context.Context, sub string) (*Token, error)
	Refresh(ctx context.Context, token string) error
	Verify(ctx context.Context, token string) (string, error)
	Clear(ctx context.Context, token string) error
}
