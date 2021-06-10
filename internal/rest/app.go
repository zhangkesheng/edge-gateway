package rest

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhangkesheng/edge-gateway/pkg/edge"
)

type App struct {
	Router   http.Handler
	stop     context.CancelFunc
	instance edge.App
}

func New() *App {
	app := &App{
	}
	app.reload()
	return app
}

func (app *App) Server() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := context.WithCancel(context.Background())
	defer stop()
	app.stop = stop

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.Router,
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

func (app *App) reload() {
	router := gin.New()

	router.LoadHTMLGlob("D:/code/mbxc/go/src/edge-gateway-github/web/*")

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	router.GET("reload", func(c *gin.Context) {
		app.reload()
		c.JSON(http.StatusOK, gin.H{"reload": true})
		app.stop()
	})

	edges := app.instance.Edges()
	for _, edge := range edges {
		edge.Router(router.Group(edge.Namespace()))
	}

	app.Router = router
}
