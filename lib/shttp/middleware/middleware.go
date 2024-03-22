package middleware

import "github.com/gin-gonic/gin"

type Middleware interface {
	Name() string
	Handle() gin.HandlerFunc
}
type ErrorMiddleware interface {
	Name() string
	Handle() gin.HandlerFunc
	Track(ctx *gin.Context, err error) error
}
