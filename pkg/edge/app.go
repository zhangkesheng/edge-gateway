package edge

import (
	"database/sql"

	"github.com/go-redis/redis"
	"github.com/zhangkesheng/edge-gateway/pkg/account"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (app *App) Edges() []Api {
	var apis []Api

	// TODO load edge
	apis = append(apis, testEdge())

	return apis
}

func testEdge() *Edge {
	return &Edge{
		Info: Info{
			Name:     "Test Edge",
			Desc:     "A test edge demo.",
			Version:  "v0.0.1",
			BasePath: "test",
		},
		AccountSvc: account.NewAccount(
			"http;//127.0.0.1:8080",
			"Test",
			"Test",
			600,
			redis.NewClient(&redis.Options{
				Addr: "",
			}),
			&sql.DB{}),
	}
}
