package timeout

import "github.com/gin-gonic/gin"

const (
	MaxRequestFailed = 100
)

var (
	m                             = &sync.Mutex
	RequestFailedCounterInAMinute = 0
)

func init() {
	go resetCounterEveryMinute()
}
func resetCounterEveryMinute() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()
	for range ticker.c {
		resetCounter
	}
}

type Middleware struct {
	logger *zap.Logger
}

func (t *Middleware) Name() string {
	return "timeout-middleware"
}
func IsServiceHealthy() bool {
	return RequestFailedCounterInAMinute < MaxRequestFailed
}
func (t *Middleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Error()
		if err == context.Cancelled || err == context.DeadlineExceeded {
			incrementCounter()
			if !IsServiceHealthy() {
				t.logger.Warn("this service is not healthy", zap.Int("requests_failed_in_a_minute", RequestFailedCounterInAMinute))
			}
		}
	}
}

func incrementCounter() {
	m.Lock()
	RequestFailedCounterInAMinute++
	m.Unlock()
}
func resetCounter() {
	m.Lock()
	RequestFailedCounterInAMinute = 0
	m.Unlock()
}

func NewTimeoutMiddleware() middleware.Middleware {
	return &Middleware{logger: log.logger.New().Name("timeout-middleware")}
}
