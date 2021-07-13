package oauth

import (
	"context"

	"github.com/zhangkesheng/edge-gateway/api/v1"
)

type MockOauth struct {
	AuthResponse         *api.AuthResponse
	AccessTokenResponse  *api.AccessTokenResponse
	RefreshTokenResponse *api.RefreshTokenResponse
	ProfileResponse      *api.ProfileResponse
}

func (m *MockOauth) Auth(context.Context, *api.AuthRequest) (*api.AuthResponse, error) {
	return m.AuthResponse, nil
}

func (m *MockOauth) AccessToken(context.Context, *api.AccessTokenRequest) (*api.AccessTokenResponse, error) {
	return m.AccessTokenResponse, nil

}

func (m *MockOauth) RefreshToken(context.Context, *api.RefreshTokenRequest) (*api.RefreshTokenResponse, error) {
	return m.RefreshTokenResponse, nil
}

func (m *MockOauth) Profile(context.Context, *api.ProfileRequest) (*api.ProfileResponse, error) {
	return m.ProfileResponse, nil
}
