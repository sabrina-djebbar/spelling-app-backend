package postgres

import (
	"context"

	"github.com/dojo-engineering/consumer-backend/lib/wtrace"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"go.opentelemetry.io/otel/attribute"
)

func Traced(db Database) Database {
	return &tracedDatabase{db: db}
}

type tracedDatabase struct {
	db Database
}

func (t *tracedDatabase) Exec(ctx context.Context, q string, params ...interface{}) (pgconn.CommandTag, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:Exec")
	span.SetAttributes(attribute.String("query", q))
	defer span.End()
	r, err := t.db.Exec(ctx, q, params...)
	if err != nil {
		span.RecordError(err)
	}
	return r, err
}
func (t *tracedDatabase) Query(ctx context.Context, q string, params ...interface{}) (pgx.Rows, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:Query")
	span.SetAttributes(attribute.String("query", q))
	defer span.End()
	r, err := t.db.Query(ctx, q, params...)
	if err != nil {
		span.RecordError(err)
	}
	return r, err
}
func (t *tracedDatabase) QueryRow(ctx context.Context, q string, params ...interface{}) pgx.Row {
	ctx, span := wtrace.New().Start(ctx, "postgres:QueryRow")
	span.SetAttributes(attribute.String("query", q))
	defer span.End()
	rawRow := t.db.QueryRow(ctx, q, params...)
	wrappedRow := &queryRowWrapper{
		inner: rawRow,
		onErr: func(err error) {
			span.RecordError(err)
		},
	}

	return wrappedRow
}
func (t *tracedDatabase) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	ctx, span := wtrace.New().Start(ctx, "postgres:SendBatch")
	defer span.End()
	tx, err := t.db.Begin(ctx)
	if err != nil {
		span.RecordError(err)
	}
	return tx.SendBatch(ctx, b)
}
func (t *tracedDatabase) Begin(ctx context.Context) (pgx.Tx, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:Begin")
	defer span.End()
	tx, err := t.db.Begin(ctx)
	if err != nil {
		span.RecordError(err)
	}
	return &tracedTransaction{tx: tx}, err
}

func (t *tracedDatabase) Close() {
	t.db.Close()
}

type queryRowWrapper struct {
	inner pgx.Row
	onErr func(err error)
}

func (r *queryRowWrapper) Scan(dest ...interface{}) error {
	err := r.inner.Scan(dest...)
	if err != nil && r.onErr != nil {
		r.onErr(err)
	}
	return err
}

type tracedTransaction struct {
	tx pgx.Tx
}

func (t *tracedTransaction) Exec(ctx context.Context, q string, params ...interface{}) (pgconn.CommandTag, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:Exec")
	span.SetAttributes(attribute.String("query", q))
	defer span.End()
	r, err := t.tx.Exec(ctx, q, params...)
	if err != nil {
		span.RecordError(err)
	}
	return r, err
}
func (t *tracedTransaction) Query(ctx context.Context, q string, params ...interface{}) (pgx.Rows, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:Query")
	span.SetAttributes(attribute.String("query", q))
	defer span.End()
	r, err := t.tx.Query(ctx, q, params...)
	if err != nil {
		span.RecordError(err)
	}
	return r, err
}
func (t *tracedTransaction) QueryRow(ctx context.Context, q string, params ...interface{}) pgx.Row {
	ctx, span := wtrace.New().Start(ctx, "postgres:QueryRow")
	span.SetAttributes(attribute.String("query", q))
	defer span.End()
	rawRow := t.tx.QueryRow(ctx, q, params...)
	wrappedRow := &queryRowWrapper{
		inner: rawRow,
		onErr: func(err error) {
			span.RecordError(err)
		},
	}

	return wrappedRow
}
func (t *tracedTransaction) Begin(ctx context.Context) (pgx.Tx, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:Begin")
	defer span.End()
	tx, err := t.tx.Begin(ctx)
	if err != nil {
		span.RecordError(err)
	}
	return &tracedTransaction{tx: tx}, err
}

func (t *tracedTransaction) Commit(ctx context.Context) error {
	ctx, span := wtrace.New().Start(ctx, "postgres:Commit")
	defer span.End()
	return t.tx.Commit(ctx)
}

func (t *tracedTransaction) Rollback(ctx context.Context) error {
	ctx, span := wtrace.New().Start(ctx, "postgres:Rollback")
	defer span.End()
	return t.tx.Rollback(ctx)
}

func (t *tracedTransaction) BeginFunc(ctx context.Context, f func(pgx.Tx) error) (err error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:Begin")
	defer span.End()
	f2 := func(ttx pgx.Tx) error {
		return f(&tracedTransaction{ttx})
	}
	return t.tx.BeginFunc(ctx, f2)

}
func (t *tracedTransaction) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:CopyFrom")
	defer span.End()
	return t.tx.CopyFrom(ctx, tableName, columnNames, rowSrc)

}
func (t *tracedTransaction) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	ctx, span := wtrace.New().Start(ctx, "postgres:SendBatch")
	defer span.End()
	return t.tx.SendBatch(ctx, b)

}
func (t *tracedTransaction) LargeObjects() pgx.LargeObjects {
	return t.tx.LargeObjects()

}
func (t *tracedTransaction) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:Prepare")
	defer span.End()
	span.SetAttributes(attribute.String("query", sql), attribute.String("name", name))
	return t.tx.Prepare(ctx, name, sql)

}
func (t *tracedTransaction) QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	ctx, span := wtrace.New().Start(ctx, "postgres:QueryFunc")
	defer span.End()
	span.SetAttributes(attribute.String("query", sql))
	ct, err := t.tx.QueryFunc(ctx, sql, args, scans, f)
	if err != nil {
		span.RecordError(err)
	}
	return ct, err
}
func (t *tracedTransaction) Conn() *pgx.Conn {
	return t.tx.Conn()
}
