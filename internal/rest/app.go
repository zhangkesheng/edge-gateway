package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangkesheng/edge-gateway/pkg/edge"
)

type App struct {
	router   http.Handler
	stop     context.CancelFunc
	instance edge.App
}

func New() (*App, error) {
	app := &App{
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
		Addr:    ":8080",
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

	edges, err := app.instance.Edges()
	if err != nil {
		return err
	}
	for _, route := range edges {
		if err := route.Router(router.Group(route.Namespace())); err != nil {
			return err
		}
	}

	app.router = router
	return nil
}
