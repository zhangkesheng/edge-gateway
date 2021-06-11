package account

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/api/v1"
	"github.com/zhangkesheng/edge-gateway/pkg/utils"
)

type Info struct {
	redirectUrl string
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
		Providers: providers,
	}, nil
}

func (app *App) Login(ctx context.Context, req *api.LoginRequest) (*api.LoginResponse, error) {
	onError := func(err error) (*api.LoginResponse, error) {
		return nil, errors.Wrap(err, "Account.Login")
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
	onError := func(err error) (*api.CallbackResponse, error) {
		return nil, errors.Wrap(err, "Account.Callback")
	}

	token, err := app.Token(ctx, &api.TokenRequest{
		GrantType:   "Bearer",
		Code:        req.GetCode(),
		State:       req.GetState(),
		ProviderKey: req.GetProviderKey(),
	})

	if err != nil {
		return onError(err)
	}

	redirect, err := url.Parse(app.config.info.redirectUrl)
	if err != nil {
		return onError(err)
	}

	redirect.Query().Add("token", token.GetAccessToken())
	return &api.CallbackResponse{
		RedirectUrl: redirect.String(),
	}, nil
}

func (app *App) Token(ctx context.Context, req *api.TokenRequest) (*api.TokenResponse, error) {
	onError := func(err error) (*api.TokenResponse, error) {
		return nil, errors.Wrap(err, "Account.Token")
	}
	provider, ok := app.config.providers[req.GetProviderKey()]
	if !ok {
		return onError(errors.New(fmt.Sprintf("Provider [%s] not found", req.GetProviderKey())))
	}

	result, err := provider.AccessToken(ctx, &api.AccessTokenRequest{
		Code:  req.GetCode(),
		State: req.GetState(),
	})
	if err != nil {
		return onError(err)
	}

	if result.GetIdentity() == nil {
		profile, err := provider.Profile(ctx, &api.ProfileRequest{
			AccessToken: result.GetToken().GetAccessToken(),
		})
		if err != nil {
			return onError(err)
		}
		result.Identity = profile.GetIdentity()
		result.Raw = profile.GetRaw()
	}

	// Get user account by `clientId` and `Response.OpenId`
	userAccount, err := app.config.storage.GetUserAccount(ctx, result.GetIdentity().GetSource(), result.GetIdentity().GetOpenId())
	if err != nil {
		return onError(err)
	}

	// New user when user account is null
	if userAccount == nil {
		// User sub
		sub := strings.ReplaceAll(uuid.New().String(), "-", "")

		userAccount = &UserAccount{
			UserSub:      sub,
			OpenId:       result.GetIdentity().GetOpenId(),
			UnionId:      result.GetIdentity().GetUnionId(),
			Nick:         result.GetIdentity().GetNick(),
			Source:       result.GetIdentity().GetSource(),
			Avatar:       result.GetIdentity().GetAvatar(),
			Email:        result.GetIdentity().GetEmail(),
			AccessToken:  result.GetToken().GetAccessToken(),
			RefreshToken: result.GetToken().GetRefreshToken(),
			Raw:          result.GetRaw(),
		}
		if result.GetToken().GetExpiresIn() > 0 {
			userAccount.ExpiredAt = time.Now().Add(time.Duration(result.GetToken().GetExpiresIn()) * time.Second).Unix()
		}

		if err = app.config.storage.SaveUserAccount(ctx, userAccount); err != nil {
			return onError(err)
		}
		if err = app.config.storage.SaveUser(ctx, &User{
			Sub:            sub,
			PrimaryAccount: userAccount.Id,
		}); err != nil {
			return onError(err)
		}
	}

	// TODO: Modify user account when login again

	// New token
	token, err := app.config.sm.New(ctx, userAccount.UserSub)
	if err != nil {
		return onError(err)
	}

	// Generate id token
	idToken, err := utils.JwtEncode(result.GetIdentity(), req.GetProviderKey())
	if err != nil {
		return onError(err)
	}
	// Create token
	return &api.TokenResponse{
		AccessToken: token.AccessToken,
		ExpiresIn:   token.ExpiresIn,
		IdToken:     idToken,
	}, nil
}

func (app *App) Refresh(ctx context.Context, req *api.RefreshRequest) (*api.RefreshResponse, error) {
	onError := func(err error) (*api.RefreshResponse, error) {
		return nil, errors.Wrap(err, "Account.Refresh")
	}
	if _, err := app.config.sm.Verify(ctx, req.GetToken()); err != nil {
		return onError(err)
	}
	token, err := app.config.sm.Refresh(ctx, req.GetToken())
	if err != nil{
		return onError(err)
	}
	return &api.RefreshResponse{
		AccessToken:token.AccessToken,
		ExpiresIn:token.ExpiresIn,
	}, nil
}

func (app *App) Verify(ctx context.Context, req *api.VerifyRequest) (*api.VerifyResponse, error) {
	onError := func(err error) (*api.VerifyResponse, error) {
		return nil, errors.Wrap(err, "Account.Verify")
	}

	if sub, err := app.config.sm.Verify(ctx, req.GetToken()); err != nil {
		return onError(err)
	} else {
		return &api.VerifyResponse{
			Sub: sub,
		}, nil
	}
}

func (app *App) Logout(ctx context.Context, req *api.LogoutRequest) (*api.LogoutResponse, error) {
	onError := func(err error) (*api.LogoutResponse, error) {
		return nil, errors.Wrap(err, "App.Logout")
	}
	if err := app.config.sm.Clear(ctx, req.GetToken()); err != nil {
		return onError(err)
	}

	return &api.LogoutResponse{}, nil
}

func New(config Config) api.AccountServer {
	return &App{config: config}
}

func NewAccount(redirectUrl, secret, issuer string, expiresIn int64, redisCli *redis.Client, db *sql.DB) api.AccountServer {
	return &App{config: Config{
		info: Info{
			redirectUrl: redirectUrl,
		},
		sm:        newRedisSessionManager(redisCli, expiresIn, secret, issuer),
		storage:   newRdsStorage(db),
		providers: map[string]api.OAuthClientServer{},
	}}

}
