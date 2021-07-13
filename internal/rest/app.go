package rest

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"github.com/zhangkesheng/edge-gateway/pkg/edge"
	"github.com/zhangkesheng/edge-gateway/pkg/oauth"
)

type App struct {
	port     string
	router   http.Handler
	stop     context.CancelFunc
	instance *edge.App
}

// TODO: 简化配置
func getOptions() []edge.Option {
	mu := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=true", "root", "", "127.0.0.1", 3306, "edge")
	db, err := sql.Open("mysql", mu)
	if err != nil {
		logrus.Fatal("Connect db error", err)
	}

	rc := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	// TODO: 测试代码, 不要提交
	return []edge.Option{
		{
			Name:               "bestzks",
			Desc:               "Bestzks api",
			Version:            "0.0.1",
			BasePath:           "bestzks",
			TokenSecret:        "bestzks123456.",
			TokenExpired:       7 * 24 * 60 * 60,
			AccountRedirectUrl: "bestzks",
			Db:                 db,
			RedisCli:           rc,
			AuthCli: []oauth.Option{
				{
					Source:          oauth.DingTalkLogin,
					ClientId:        "dingoa7thvgz9fycf710ys",
					Secret:          "8UrY9ZAwbNpFeedQ67A8nrWyoeRmBbEezi2Xe9hzd_4lRpM2WI8wQI9GWgr-UZeG",
					DefaultRedirect: "https://auth.bestzks.com/app?clientId=dingoa7thvgz9fycf710ys",
				},
			},
			Backends: []edge.BackendOption{
				{
					BathPath: "test",
					Host:     "127.0.0.1:80",
					Apis:     []string{},
				},
			},
		},
	}
}

func New() (*App, error) {
	instance, err := edge.New(getOptions())
	if err != nil {
		return nil, err
	}
	app := &App{
		instance: instance,
		port:     ":8080",
	}
	if err := app.reload(); err != nil {
		return nil, err
	}
	return app, nil
}

func (app *App) Server() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	app.stop = stop

	srv := &http.Server{
		Addr:    app.port,
		Handler: app.router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()
	if err := srv.Shutdown(context.Background()); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Shutdown: %s\n", err)
	}

	app.Server()
}

func (app *App) reload() error {
	router := gin.New()

	router.LoadHTMLGlob("D:/code/mbxc/go/src/edge-gateway-github/web/*")

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("reload", func(c *gin.Context) {
		err := app.reload()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"reload": true})
		app.stop()
	})

	edges := app.instance.Edges()
	for _, route := range edges {
		if err := route.Router(router.Group(route.Namespace())); err != nil {
			return err
		}
	}

	app.router = router
	return nil
}
