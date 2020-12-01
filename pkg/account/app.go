package account

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

type Info struct {
	Title   string
	Desc    string
	Favicon string
}

type Config struct {
	info      Info
	sm        SessionManager
	storage   Storage
	providers map[string]api.OAuthClientServer
}

type App struct {
	config Config
}

func (app *App) Info(ctx context.Context, req *empty.Empty) (*api.InfoResponse, error) {
	var providers []*api.InfoResponse_Provider
	for v, _ := range app.config.providers {
		providers = append(providers, &api.InfoResponse_Provider{
			Type: v,
			Key:  v,
		})
	}
	return &api.InfoResponse{
		Name:      app.config.info.Title,
		Desc:      app.config.info.Desc,
		Providers: providers,
	}, nil
}

func (app *App) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	onError := func(err error) (*api.LoginResponse, error) {
		return nil, errors.Wrap(err, "App.Login")
	}

	client, ok := app.config.providers[req.GetProviderKey()]
	if !ok {
		return onError(errors.New(fmt.Sprintf("Provider [%s] not found", req.GetProviderKey())))
	}

	authReq := &api.AuthRequest{
		Scope:        req.GetScope(),
		ResponseType: req.GetResponseType(),
		RedirectUrl:  req.GetRedirectUrl(),
		State:        req.GetState(),
	}

	resp, err := client.Auth(ctx, authReq)
	if err != nil {
		return onError(err)
	}

	return &api.LoginResponse{
		RedirectTo: resp.GetRedirectTo(),
	}, nil
}

func (app *App) Callback(ctx context.Context, req *api.CallbackRequest) (*api.CallbackResponse, error) {
	panic("implement me")
}

func (app *App) Token(ctx context.Context, req *api.TokenRequest) (*api.TokenResponse, error) {
	panic("implement me")
}

func (app *App) Refresh(ctx context.Context, req *api.RefreshRequest) (*api.RefreshResponse, error) {
	panic("implement me")
}

func (app *App) Verify(ctx context.Context, req *api.VerifyRequest) (*api.VerifyResponse, error) {
	panic("implement me")
}

func (app *App) Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutResponse, error) {
	panic("implement me")
}

func New(config Config) api.AccountServer {
	return &App{config: config}
}
