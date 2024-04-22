package shttp

import (
	"fmt"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/shttp/middleware"
	"log"
)

var logger = log.Logger{}

func (srv *server) RegisterMiddleware(middleware middleware.Middleware) {
	fmt.Println("Registering Middleware %s", middleware.Name())

	srv.router.Use(middleware.Handle())
}
