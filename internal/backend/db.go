package backend

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"entgo.io/ent/dialect"
	"github.com/ProtonMail/gluon/internal/backend/ent"
)

type DB struct {
	db   *ent.Client
	lock sync.RWMutex
}

func (d *DB) Init(ctx context.Context) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	return d.db.Schema.Create(ctx)
}

func (d *DB) Read(ctx context.Context, fn func(context.Context, *ent.Client) error) error {
	_, err := DBReadResult(ctx, d, func(ctx context.Context, client *ent.Client) (struct{}, error) {
		return struct{}{}, fn(ctx, client)
	})

	return err
}

func (d *DB) Write(ctx context.Context, fn func(context.Context, *ent.Tx) error) error {
	_, err := DBWriteResult(ctx, d, func(ctx context.Context, tx *ent.Tx) (struct{}, error) {
		return struct{}{}, fn(ctx, tx)
	})

	return err
}

func (d *DB) Close() error {
	return d.db.Close()
}

func DBReadResult[T any](ctx context.Context, db *DB, fn func(context.Context, *ent.Client) (T, error)) (T, error) {
	db.lock.RLock()
	defer db.lock.RUnlock()

	return fn(ctx, db.db)
}

func DBWriteResult[T any](ctx context.Context, db *DB, fn func(context.Context, *ent.Tx) (T, error)) (T, error) {
	db.lock.Lock()
	defer db.lock.Unlock()

	var failResult T

	tx, err := db.db.Tx(ctx)
	if err != nil {
		return failResult, err
	}

	defer func() {
		if v := recover(); v != nil {
			if err := tx.Rollback(); err != nil {
				panic(fmt.Errorf("rolling back while recovering (%v): %w", v, err))
			}

			panic(v)
		}
	}()

	result, err := fn(ctx, tx)
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			return failResult, fmt.Errorf("rolling back transaction: %w", rerr)
		}

		return failResult, err
	}

	if err := tx.Commit(); err != nil {
		return failResult, fmt.Errorf("committing transaction: %w", err)
	}

	return result, nil
}

func (b *Backend) newDB(userID string) (*DB, error) {
	dir := filepath.Join(b.dir, "db")

	if err := os.MkdirAll(dir, 0o700); err != nil {
		return nil, err
	}

	client, err := ent.Open(dialect.SQLite, getDatabasePath(dir, userID))
	if err != nil {
		return nil, err
	}

	return &DB{db: client}, nil
}

func getDatabasePath(dir, userID string) string {
	return fmt.Sprintf("file:%v?cache=shared&_fk=1", filepath.Join(dir, fmt.Sprintf("%v.db", userID)))
}