package shttp

import (
	"context"
	"github.com/gin-gonic/gin"
)

//type ctxKey string

//const RequestContextKey ctxKey = "http_request"

func createContext(gctx *gin.Context) context.Context {
	ctx := context.WithValue(gctx, "http_request", gctx.Request)
	for k, v := range gctx.Keys {
		ctx = context.WithValue(ctx, k, v)
	}
	return ctx
}

/*func GetRequestFromContext(ctx context.Context) *http.Request {
	value := ctx.Value(RequestContextKey)
	if request, ok := value.(*http.Request); ok {
		return request
	}
	return nil
}*/
