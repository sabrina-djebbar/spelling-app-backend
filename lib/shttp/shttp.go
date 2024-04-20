package shttp

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/id"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"time"
)

type HTTPServer interface {
	RegisterHandler(path string, handler interface{})
	RegisterGinHandler(method string, path string, handler gin.HandlerFunc)
	Listen(addr string) error
}

type server struct {
	router       *gin.Engine
	http         *http.Server
	errorTracker func(ctx *gin.Context, err interface{}) error
}

var HealthPath = "/"

func New(cmd *cobra.Command) HTTPServer {
	router := gin.New()

	router.GET(HealthPath, func(ctx *gin.Context) { ctx.JSON(200, gin.H{"ready": "ok"}) })

	router.Use(NewServiceIdentityMiddleware(cmd))
	return &server{router: router}
}

func (srv *server) Listen(port string) error {
	envPort := os.Getenv("PORT")
	if envPort != "" {
		port = envPort
	}
	srv.http = &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		ReadTimeout:  15 * time.Second, // max time to read request from the client
		WriteTimeout: 15 * time.Second, // max time to write response to the client
		IdleTimeout:  60 * time.Second,
		Handler:      srv.router,
	}
	log.Printf("listening on %s\n", srv.http.Addr)
	// killable.RegisterKillable(srv.Close)
	if err := srv.http.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
	return nil
}
func (srv *server) Close(ctx context.Context) {
	if err := srv.http.Shutdown(ctx); err != nil {
		panic(err)
	}
}

var ServedBy = ""

func NewServiceIdentityMiddleware(cmd *cobra.Command) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// request header
		requestID := ctx.GetString("http_request_identifier")
		if requestID == "" {
			requestID = ctx.GetHeader("X-Request-ID")
			if requestID == "" {
				requestID = id.Generate("http_request")
			}
			ctx.Set("http_request_identifier", requestID)
		}
		ctx.Header("X-Request-ID", requestID)

		if ServedBy != "" {
			ctx.Header("X-ServedBy", ServedBy)
		}

		ctx.Header("X-Served-Date", time.Now().UTC().Format(time.RFC3339))
		ctx.Next()
	}
}
