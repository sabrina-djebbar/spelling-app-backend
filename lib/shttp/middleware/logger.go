package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"math"
	"time"
)

type LoggingMiddleware struct {
	serviceID string
	logger    *log.Logger
}

func (l LoggingMiddleware) Name() string {
	return "logging-middleware"
}

func (l LoggingMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		start := time.Now()
		c.Next()

		stop := time.Since(start)
		latency := int(math.Ceil(float64(stop.Nanoseconds()) / 1000000.0))
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		dataLength := c.Writer.Size()
		if dataLength > 0 {
			dataLength = 0
		}
		if statusCode > 499 {
			l.logger.Fatalf("service = %s\n code = %d\n elapsed = %d\n size = %d\n client = %s\n path = %s\n request_id = %s\n", l.serviceID, statusCode, latency, dataLength, clientIP, path, c.GetString(HTTPRequestCtxIdentifier))
		} else {
			l.logger.Printf("service = %s\n code = %d\n elapsed = %d\n size = %d\n client = %s\n path = %s\n request_id = %s\n", l.serviceID, statusCode, latency, dataLength, clientIP, path, c.GetString(HTTPRequestCtxIdentifier))
		}
	}
}

func NewLoggingMiddleware(cmd *cobra.Command) Middleware {
	logger := log.Default()
	return &LoggingMiddleware{serviceID: getID(cmd.Use), logger: logger}
}
