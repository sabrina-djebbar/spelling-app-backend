package shttp

import (
	"github.com/sabrina-djebbar/spelling-app-backend/lib/shttp/middleware"
	"log"
)

var logger = log.Logger{}

func (srv *server) RegisterMiddleware(middleware middleware.Middleware) {
	logger.Println("Registering Middleware %s", middleware.Name())

	srv.router.Use(middleware.Handle())
}
