package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	idGen "github.com/sabrina-djebbar/spelling-app-backend/lib/id"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"time"
)

const (
	// HTTPRequestCtxIdentifier is the key for the request id in the gin context
	HTTPRequestCtxIdentifier = "http_request_identifier"
	// HTTPRequestHeaderIdentifier is the key for the request id in the http headers
	HTTPRequestHeaderIdentifier = "X-Request-ID"

	requestResourceIdentifier = "http_request"
)

var ServedBy = ""

func NewServiceIdentityMiddleware(cmd *cobra.Command) gin.HandlerFunc {
	if cmd != nil {
		ServedBy = getID(cmd.Use)
	}

	return func(ctx *gin.Context) {
		// Set our request identifier header
		setRequestIdentifier(ctx)

		// Set our served by header
		if ServedBy != "" {
			ctx.Header("X-Served-By", ServedBy)
		}

		// Add the current time as a header.
		// This allows the app to pull it and detect if the current time is out of sync.
		// We put our own in place rather than using Date as both the LB and Cloudflare will override this.
		ctx.Header("X-Served-Date", time.Now().UTC().Format(time.RFC3339))

		ctx.Next()
	}
}

func setRequestIdentifier(ctx *gin.Context) {
	requestID := ctx.GetString(HTTPRequestCtxIdentifier)
	if requestID == "" {
		requestID = ctx.GetHeader(HTTPRequestHeaderIdentifier)
		if requestID == "" {
			requestID = idGen.Generate(requestResourceIdentifier)
		}

		ctx.Set(HTTPRequestCtxIdentifier, requestID)
	}

	ctx.Header(HTTPRequestHeaderIdentifier, requestID)
}

func getID(serviceName string) string {
	id, err := os.Hostname()
	if err != nil {
		return serviceName
	}

	// Ignore hostnames that have a dot in them.
	// They are not a docker container in this case.
	if strings.Contains(id, ".") {
		return serviceName
	}

	// If the hostname begins with our service name then don't duplicate
	if strings.HasPrefix(id, serviceName) {
		return id
	}

	return fmt.Sprintf("%s-%s", serviceName, id)
}
