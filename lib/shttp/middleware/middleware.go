package middleware

import "github.com/gin-gonic/gin"

// Middleware is our wrapper above a gin.HandlerFunc
type Middleware interface {
	Name() string

	Handle() gin.HandlerFunc
}

// ErrorMiddleware is a special type of middleware that allows tracking of serr in non fatal crashes
type ErrorMiddleware interface {
	Name() string

	Handle() gin.HandlerFunc
	Track(ctx *gin.Context, err error) error
}
