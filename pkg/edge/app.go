package edge

import (
	"database/sql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/zhangkesheng/edge-gateway/api/v1"
	"github.com/zhangkesheng/edge-gateway/pkg/oauth"
)

type App struct {
	apis []Api
}

func New() *App {
	app := &App{}
	app.Reload()
	return app
}

func (app *App) Edges() []Api {
	return app.apis
}

type Options struct {
	Name        string
	Desc        string
	Version     string
	BasePath    string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPws       string
	DbDatabase  string
	RedisHost   string
	RedisPort   string
	RedisPwd    string
	TokenSecret string
	// ms
	TokenExpired       int64
	AccountRedirectUrl string
	AuthCli            []AuthOption
}
type AuthOption struct {
	Type            oauth.Source
	ClientId        string
	Secret          string
	AuthUrl         string
	LogoutUrl       string
	AccessTokenUrl  string
	ApiUrl          string
	DefaultRedirect string
	DefaultScope    string
}

func (app *App) Reload() error {
	onError := func(err error) error {
		return errors.Wrap(err, "Edge reload")
	}
	// TODO read config from file. like: yaml or json or db.
	var options []Options

	for _, option := range options {
		// TODO check option
		// TODO 支持其他的存储方式
		mu := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", option.DbUser, option.DbPws, option.DbHost, option.DbPort, option.DbDatabase)
		db, err := sql.Open("mysql", mu)
		if err != nil {
			return onError(err)
		}

		redisCli := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", option.RedisHost, option.RedisPort),
			Password: option.RedisPwd,
		})

		edgeConfig := Config{
			Info: Info{
				Name:     option.Name,
				Desc:     option.Desc,
				Version:  option.Version,
				BasePath: option.BasePath,
			},
			DB:                 db,
			RedisCli:           redisCli,
			AccountRedirectUrl: option.AccountRedirectUrl,
			TokenSecret:        option.TokenSecret,
			TokenExpired:       option.TokenExpired,
			AuthClient:         map[string]api.OAuthClientServer{},
		}

		for _, cli := range option.AuthCli {
			edgeConfig.AuthClient[cli.ClientId] = oauth.New(cli.Type, oauth.Config{
				ClientId:        cli.ClientId,
				Secret:          cli.Secret,
				AuthUrl:         cli.AuthUrl,
				LogoutUrl:       cli.LogoutUrl,
				AccessTokenUrl:  cli.AccessTokenUrl,
				ApiUrl:          cli.ApiUrl,
				DefaultRedirect: cli.DefaultRedirect,
				DefaultScope:    cli.DefaultScope,
			})
		}
		app.apis = append(app.apis, NewEdge(edgeConfig))
	}

	return nil
}
