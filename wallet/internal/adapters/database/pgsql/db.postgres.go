package postgresdb

import (
	"context"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	port "github.com/kaolnwza/muniverse/wallet/internal/ports"
	_ "github.com/lib/pq"
)

func NewPostgresTransactor(db *sqlx.DB) port.Transactor {
	return &postgresRepo{conn: db}
}

func NewPostgresDB(url string, dbName string) *sqlx.DB {
	conn, err := sqlx.Open("postgres", url)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Successfully | " + dbName + " | Database Connected!")

	return conn
}

type txKey struct{}

func injectTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, txKey{}, tx)
}

func extractTx(ctx context.Context) *sqlx.Tx {
	if tx, ok := ctx.Value(txKey{}).(*sqlx.Tx); ok {
		return tx
	}
	return nil
}

func (db *postgresRepo) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.SelectContext(ctx, dest, query, args...)
	}

	return db.conn.SelectContext(ctx, dest, query, args...)
}

func (db *postgresRepo) GetOne(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.GetContext(ctx, dest, query, args...)
	}

	return db.conn.GetContext(ctx, dest, query, args...)
}

func (db *postgresRepo) Update(ctx context.Context, query string, args ...interface{}) error {
	tx := extractTx(ctx)
	if tx != nil {
		_, err := tx.ExecContext(ctx, query, args...)
		return err
	}

	_, err := db.conn.ExecContext(ctx, query, args...)
	return err
}

func (db *postgresRepo) Insert(ctx context.Context, query string, args ...interface{}) error {
	tx := extractTx(ctx)
	if tx != nil {
		_, err := tx.ExecContext(ctx, query, args...)
		return err
	}

	_, err := db.conn.ExecContext(ctx, query, args...)
	return err
}

func (db *postgresRepo) InsertWithReturning(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.SelectContext(ctx, dest, query, args...)
	}

	return db.conn.SelectContext(ctx, dest, query, args...)
}

func (db *postgresRepo) InsertWithReturningOne(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	tx := extractTx(ctx)
	if tx != nil {
		return tx.GetContext(ctx, dest, query, args...)
	}

	return db.conn.GetContext(ctx, dest, query, args...)
}

type postgresRepo struct {
	conn *sqlx.DB
}

func NewPostgresRepo(db *sqlx.DB) port.Transactor {
	return &postgresRepo{conn: db}
}

func (db *postgresRepo) WithinTransaction(ctx context.Context, tFunc func(ctx context.Context) error) error {
	tx, err := db.conn.Beginx()
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	fmt.Println("-------Postgres Transaction: Begin---------")

	err = tFunc(injectTx(ctx, tx))
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Printf("rollback transaction: %v", errRollback)
		}

		fmt.Println("-------Postgres Transaction: Rollback-------")
		return err
	}

	if errCommit := tx.Commit(); errCommit != nil {
		log.Printf("commit transaction: %v", errCommit)
	}

	fmt.Println("-------Postgres Transaction: Committed------")
	return nil
}
