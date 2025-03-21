package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/jackc/pgx/v5/stdlib"
)


type SqlGo struct {
	dB *sql.DB
}

func (db *SqlGo) Open(databaseConfig DatabaseConfig) error {
	// Create DSN for pgx
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		databaseConfig.User, databaseConfig.Password, databaseConfig.Host, databaseConfig.Port, databaseConfig.Database)

	// Open using pgx
	var err error
	db.dB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Check connection
	err = db.dB.Ping()
	if err != nil {
		db.dB.Close()
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Database connected successfully!")
	return nil
}

func (db *SqlGo) Close() error {
	if db.dB != nil {
		return db.dB.Close()
	}
	return nil
}


func (db *SqlGo) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.dB.QueryContext(ctx, query, args...)
}

func (db *SqlGo) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.dB.ExecContext(ctx, query, args...)
}

func (db *SqlGo) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	tx, err := db.dB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &SqlTx{tx: tx}, nil
}


type SqlTx struct {
	tx *sql.Tx
}

func (tx *SqlTx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, query, args...)
}

func (tx *SqlTx) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.tx.QueryContext(ctx, query, args...)
}

func (tx *SqlTx) Commit() error {
	return tx.tx.Commit()
}

func (tx *SqlTx) Rollback() error {
	return tx.tx.Rollback()
}
