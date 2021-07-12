package main

import (
	"log"

	"github.com/zhangkesheng/edge-gateway/internal/rest"
)

func main() {
	app, err := rest.New()
	if err != nil {
		log.Fatal("New app error.", err)
	}
	app.Server()
}
