package account

import (
	"context"
	"strings"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/assert"
	"github.com/zhangkesheng/edge-gateway/api/v1"
	"github.com/zhangkesheng/edge-gateway/pkg/oauth"
)

func TestAccount(t *testing.T) {
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()
	redisCli := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	sm := newRedisSessionManager(redisCli, 60*1000, "Test", "Test")

	app := &App{
		info: Info{
			basePath:    "test",
			name:        "test",
			desc:        "test",
			redirectUrl: "test",
		},
		storage: newMockStorage(),
		sm:      sm,
		providers: map[string]*oauthCli{
			"test": {cli: &oauth.MockOauth{
				AuthResponse: &api.AuthResponse{
					RedirectTo: "http://127.0.0.1:8080/test?response_type=code&app_id=test",
				},
				AccessTokenResponse: &api.AccessTokenResponse{
					Token: &api.Token{
						AccessToken: "TestToken",
					},
					Identity: nil,
					Raw:      "",
				},
			}, source: "test"},
		},
	}

	ctx := context.Background()
	t.Run("info", func(t *testing.T) {
		resp, err := app.Info(ctx, &empty.Empty{})
		if assert.NoError(t, err) {
			assert.Equal(t, 1, len(resp.Providers))
		}
	})

	t.Run("login", func(t *testing.T) {
		resp, err := app.Login(ctx, &api.LoginRequest{
			ProviderKey: "test",
		})
		if assert.NoError(t, err) {
			assert.Equal(t, "http://127.0.0.1:8080/test?response_type=code&app_id=test", resp.GetRedirectTo())
		}
	})

	t.Run("callback", func(t *testing.T) {
		resp, err := app.Callback(ctx, &api.CallbackRequest{
			Code:        "t123",
			State:       "",
			ProviderKey: "test",
		})

		if assert.NoError(t, err) {
			assert.True(t, strings.Index(resp.GetRedirectUrl(), "?token=") > 0)
		}
	})
}
