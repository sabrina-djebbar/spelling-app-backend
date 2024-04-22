package postgres

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/IBM/pgxpoolprometheus"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/dojo-engineering/consumer-backend/lib/killable"
	"github.com/dojo-engineering/consumer-backend/lib/werr"

	"cloud.google.com/go/cloudsqlconn"
)

var defaultMaxOpenConns = 15
var defaultMaxIdleConns = 15
var defaultConnMaxLifetime = time.Hour
var defaultConnMaxIdleTime = time.Minute * 5

type Middleware func(Database) Database

type Config struct {
	DSN         string `json:"dsn" env:"DSN,required"`
	Instance    string `json:"instance" env:"INSTANCE,required"`
	SkipMetrics bool   `json:"-" env:"SKIP_METRICS, default=false"`
}

var ErrNoRows = pgx.ErrNoRows

// New returns a traced Database through a pgx connection pool
func New(ctx context.Context, name string, cfg Config) (Database, error) {
	connPool, err := CreatePgxPool(ctx, name, cfg)
	if err != nil {
		return nil, err
	}
	if !cfg.SkipMetrics {
		collector := pgxpoolprometheus.NewCollector(connPool, map[string]string{"db_name": name})
		prometheus.MustRegister(collector)
	}

	return Traced(connPool), nil
}

func NewWithoutTracing(ctx context.Context, name string, cfg Config) (Database, error) {
	db, err := CreatePgxPool(ctx, name, cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreatePgxPool returns a new pgx pool to expose access to lower level apis
// Main use for data migration CLIs
func CreatePgxPool(ctx context.Context, name string, cfg Config) (*pgxpool.Pool, error) {
	pgxconfig, err := pgxpool.ParseConfig(cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to parse pgx config: %w", err)
	}

	pgxconfig.MaxConns = int32(defaultMaxOpenConns)
	pgxconfig.MaxConnIdleTime = defaultConnMaxIdleTime
	pgxconfig.MaxConnLifetime = defaultConnMaxLifetime

	// Use IAM auth only when instance value is provided, otherwise user/pwd approach is used
	if cfg.Instance != "" {
		// Create a new dialer with IAM auth
		d, err := cloudsqlconn.NewDialer(ctx, cloudsqlconn.WithIAMAuthN())
		if err != nil {
			return nil, fmt.Errorf("failed to initialize dialer: %w", err)
		}

		killable.RegisterKillable(func(_ context.Context) {
			d.Close()
		})

		// Tell the driver to use the Cloud SQL Go Connector to create connections
		pgxconfig.ConnConfig.DialFunc = func(ctx context.Context, _ string, instance string) (net.Conn, error) {
			return d.Dial(ctx, cfg.Instance)
		}
	}

	connPool, err := pgxpool.ConnectConfig(ctx, pgxconfig)
	if err != nil {
		return nil, werr.Wrap(err, werr.WithMessage("unable to open database"))
	}

	return connPool, nil
}
