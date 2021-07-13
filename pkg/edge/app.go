package edge

import (
	"database/sql"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/pkg/account"
	"github.com/zhangkesheng/edge-gateway/pkg/backend"
	"github.com/zhangkesheng/edge-gateway/pkg/oauth"
	"github.com/zhangkesheng/edge-gateway/pkg/types"
)

type App struct {
	options []Option
	apis    []types.ApiRoute
}

func New(options []Option) (*App, error) {
	app := &App{
		options: options,
	}
	if err := app.Reload(); err != nil {
		return nil, errors.Wrap(err, "NewOauthCli edge")
	}
	return app, nil
}

func (app *App) Edges() []types.ApiRoute {
	return app.apis
}

type Option struct {
	Name        string
	Desc        string
	Version     string
	BasePath    string
	TokenSecret string
	// ms
	TokenExpired       int64
	AccountRedirectUrl string
	Db                 *sql.DB
	RedisCli           *redis.Client
	AuthCli            []oauth.Option
	Backends           []BackendOption
}

// TODO: use open-api
type BackendOption struct {
	BathPath string
	Host     string
	Apis     []string
}

func (app *App) Reload() error {
	// onError := func(err error) error {
	// 	return errors.Wrap(err, "Edge reload")
	// }
	// TODO read config from file. like: yaml or json or db.

	app.apis = []types.ApiRoute{}

	for _, option := range app.options {
		// TODO check option
		var accountProviders []oauth.Option
		for _, cli := range option.AuthCli {
			accountProviders = append(accountProviders, oauth.Option{
				Source:          cli.Source,
				ClientId:        cli.ClientId,
				Secret:          cli.Secret,
				DefaultRedirect: cli.DefaultRedirect,
			})
		}

		accountSvc := account.New(account.Option{
			Name:        option.Name,
			Desc:        option.Desc,
			RedirectUrl: option.AccountRedirectUrl,
			Secret:      option.TokenSecret,
			Issuer:      option.Name,
			BasePath:    option.BasePath,
			ExpiresIn:   option.TokenExpired,
			RedisCli:    option.RedisCli,
			Db:          option.Db,
			Providers:   accountProviders,
		})

		backendMap := map[string]types.ApiRoute{}
		for _, b := range option.Backends {
			backendMap[b.BathPath] = backend.NewReverseProxy(b.Host, b.Apis)
		}

		app.apis = append(app.apis, &Edge{
			Name:       option.Name,
			Desc:       option.Desc,
			Version:    option.Version,
			BasePath:   option.BasePath,
			AccountSvc: accountSvc,
		})
	}

	return nil
}
