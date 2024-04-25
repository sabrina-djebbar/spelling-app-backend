package killable

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"go.uber.org/atomic"

	"golang.org/x/sync/errgroup"
)

// Killhooks should begin the shutdown
var m = sync.Mutex{}
var killableServices = make([]func(ctx context.Context), 0)

// Warnings should signal that the service should prepare to shutdown
var killWarning = atomic.NewBool(false)

// The default kubernetes shutdown grace is 30 seconds, so ideally the two below add up to less than 30.
const warnDeadline = time.Second * 20
const killDeadline = time.Second * 5

// RegisterKillable will allow you to register a function that will be called on the sigterm request from the client
func RegisterKillable(kill func(ctx context.Context)) {
	// Ensure multiple calls to this function at once do not loose killableServices
	m.Lock()
	defer m.Unlock()

	killableServices = append(killableServices, kill)
}

// IsKillImpending returns true when a kill is upcoming. This allows services to be warned.
func IsKillImpending() bool {
	return killWarning.Load()
}

func handleWarns() {
	killWarning.Store(true)
	time.Sleep(warnDeadline)
}

func handleKills() {
	// Ensure no new killableServices can be registered during shutdown
	m.Lock()

	ctx, cancel := context.WithTimeout(context.Background(), killDeadline)
	defer cancel()

	g, gctx := errgroup.WithContext(ctx)

	for i := range killableServices {
		killable := killableServices[i]

		g.Go(func() error {
			killable(gctx)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		panic(err)
	}
}

// ListenToKill will listen to a kill request before closing
func ListenToKill() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	<-c // Wait for sig

	handleWarns()
	handleKills()

	os.Exit(0)
}
