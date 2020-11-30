package account

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/zhangkesheng/edge-gateway/api/v1"
)

type Info struct {
	Title   string
	Desc    string
	Favicon string
}

type Config struct {
	info    Info
	sm      SessionManager
	storage Storage
	clients map[string]api.OAuthClientServer
}

type App struct {
	config Config
}

func (app *App) Info(ctx context.Context, req *empty.Empty) (*api.InfoResponse, error) {
	var clients []*api.InfoResponse_Client
	for v, _ := range app.config.clients {
		clients = append(clients, &api.InfoResponse_Client{
			Type: v,
			Key:  v,
		})
	}
	return &api.InfoResponse{
		Name:    app.config.info.Title,
		Desc:    app.config.info.Desc,
		Clients: clients,
	}, nil
}

func (app *App) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	panic("implement me")
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
