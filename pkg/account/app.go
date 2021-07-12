package account

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/api/v1"
	"github.com/zhangkesheng/edge-gateway/pkg/types"
	"github.com/zhangkesheng/edge-gateway/pkg/oauth"
	"github.com/zhangkesheng/edge-gateway/pkg/utils"
)

const (
	loginHtml = "login.html"
)

type Info struct {
	redirectUrl string
	basePath    string
	name        string
	desc        string
}

type App struct {
	info      Info
	sm        SessionManager
	storage   Storage
	providers map[string]api.OAuthClientServer
}

func (app *App) Info(ctx context.Context, req *empty.Empty) (*api.InfoResponse, error) {
	var providers []*api.InfoResponse_Provider
	for v, _ := range app.providers {
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

	client, ok := app.providers[req.GetProviderKey()]
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

	redirect, err := url.Parse(app.info.redirectUrl)
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
	provider, ok := app.providers[req.GetProviderKey()]
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
	userAccount, err := app.storage.GetUserAccount(ctx, result.GetIdentity().GetSource(), result.GetIdentity().GetOpenId())
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

		if err = app.storage.SaveUserAccount(ctx, userAccount); err != nil {
			return onError(err)
		}
		if err = app.storage.SaveUser(ctx, &User{
			Sub:            sub,
			PrimaryAccount: userAccount.Id,
		}); err != nil {
			return onError(err)
		}
	}

	// TODO: Modify user account when login again

	// New token
	token, err := app.sm.New(ctx, userAccount.UserSub)
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
	if _, err := app.sm.Verify(ctx, req.GetToken()); err != nil {
		return onError(err)
	}
	token, err := app.sm.Refresh(ctx, req.GetToken())
	if err != nil {
		return onError(err)
	}
	return &api.RefreshResponse{
		AccessToken: token.AccessToken,
		ExpiresIn:   token.ExpiresIn,
	}, nil
}

func (app *App) Verify(ctx context.Context, req *api.VerifyRequest) (*api.VerifyResponse, error) {
	onError := func(err error) (*api.VerifyResponse, error) {
		return nil, errors.Wrap(err, "Account.Verify")
	}

	if sub, err := app.sm.Verify(ctx, req.GetToken()); err != nil {
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
	if err := app.sm.Clear(ctx, req.GetToken()); err != nil {
		return onError(err)
	}

	return &api.LogoutResponse{}, nil
}

func (app *App) Router(r gin.IRouter) error {
	acGroup := r.Group("account")
	// Login page
	acGroup.GET(loginHtml, func(c *gin.Context) {
		ctx := c.Request.Context()
		info, err := app.Info(ctx, &empty.Empty{})
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.HTML(http.StatusOK, loginHtml, gin.H{
			"basePath": app.info.basePath,
			"name":     app.info.name,
			"desc":     app.info.desc,
			"info":     info,
		})
	})
	// Info api
	acGroup.GET("", func(c *gin.Context) {
		ctx := c.Request.Context()
		info, err := app.Info(ctx, &empty.Empty{})
		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"basePath": app.info.basePath,
			"name":     app.info.name,
			"desc":     app.info.desc,
			"info":     info,
		})
	})

	// Logout
	acGroup.GET("logout", func(c *gin.Context) {
		token, err := utils.CheckToken(c)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx := c.Request.Context()
		resp, err := app.Logout(ctx, &api.LogoutRequest{
			Token: token,
		})
		utils.HandleJsonResp(c, err, resp)
	})

	// Refresh token
	acGroup.POST("refresh", func(c *gin.Context) {
		token, err := utils.CheckToken(c)
		if err != nil {
			_ = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx := c.Request.Context()
		newToken, err := app.Refresh(ctx, &api.RefreshRequest{
			Token: token,
		})
		utils.HandleJsonResp(c, err, newToken)
	})

	// Account auth client api
	acCliGroup := acGroup.Group("/client/:clientId")
	acCliGroup.GET("authorize", func(c *gin.Context) {
		clientId := c.Param("clientId")
		var req struct {
			State        string `form:"state"`
			RedirectUrl  string `form:"redirectUrl"`
			Redirect     bool   `form:"redirect"`
			ResponseType string `form:"responseType"`
			Scope        string `form:"scope"`
		}
		if err := c.BindQuery(&req); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		ctx := c.Request.Context()
		resp, err := app.Login(ctx, &api.LoginRequest{
			ResponseType: req.ResponseType,
			ProviderKey:  clientId,
			RedirectUrl:  req.RedirectUrl,
			Scope:        req.Scope,
			State:        req.State,
		})

		if err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
		}

		if req.Redirect {
			c.Redirect(http.StatusSeeOther, resp.GetRedirectTo())
		}

		utils.HandleJsonResp(c, err, resp)
	})

	acCliGroup.GET("callback", func(c *gin.Context) {
		clientId := c.Param("clientId")

		ctx := c.Request.Context()
		resp, err := app.Callback(ctx, &api.CallbackRequest{
			State:       c.Query("state"),
			Code:        c.Query("code"),
			ProviderKey: clientId,
		})

		utils.HandleJsonResp(c, err, resp)
	})

	return nil
}
func (app *App) Namespace() string {
	return "account"
}

type Option struct {
	Name, Desc, RedirectUrl, Secret, Issuer string
	ExpiresIn                               int64
	RedisCli                                *redis.Client
	Db                                      *sql.DB
	Providers                               []oauth.Option
}

func New(option Option) types.ApiRoute {
	accountSvc := &App{
		info: Info{
			redirectUrl: option.RedirectUrl,
			name:        option.Name,
			desc:        option.Desc,
		},
		sm:        newRedisSessionManager(option.RedisCli, option.ExpiresIn, option.Secret, option.Issuer),
		storage:   newRdsStorage(option.Db),
		providers: map[string]api.OAuthClientServer{},
	}

	for _, provider := range option.Providers {
		accountSvc.providers[provider.Config.ClientId] = oauth.NewOauth(provider)
	}

	return accountSvc
}
