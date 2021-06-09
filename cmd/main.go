package main

import (
	"log"
	"net/http"

	"github.com/zhangkesheng/edge-gateway/internal/rest"
)

func main() {
	app := rest.New()
	app.Server()
	srv := &http.Server{
		Addr: ":8090",
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}

func server() {

}
