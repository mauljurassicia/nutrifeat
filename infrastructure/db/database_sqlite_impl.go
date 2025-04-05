package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/ncruces/go-sqlite3"
)


type SqliteGo struct {
	dB *sql.DB
}



func (db *SqliteGo) Open(databaseConfig DatabaseConfig) error {
	// Create DSN for sqlite3
	dsn := databaseConfig.Database

	// Open using sqlite3
	var err error
	db.dB, err = sql.Open("sqlite3", dsn)
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

func (db *SqliteGo) Close() error {
	if db.dB != nil {
		return db.dB.Close()
	}
	return nil
}


func (db *SqliteGo) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.dB.QueryContext(ctx, query, args...)
}

func (db *SqliteGo) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.dB.ExecContext(ctx, query, args...)
}

func (db *SqliteGo) BeginTx(ctx context.Context, opts *sql.TxOptions) (Tx, error) {
	tx, err := db.dB.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}
	return &SqlTx{tx: tx}, nil
}


type SqliteTx struct {
	tx *sql.Tx
}

func (tx *SqliteTx) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.tx.ExecContext(ctx, query, args...)
}

func (tx *SqliteTx) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return tx.tx.QueryContext(ctx, query, args...)
}

func (tx *SqliteTx) Commit() error {
	return tx.tx.Commit()
}

func (tx *SqliteTx) Rollback() error {
	return tx.tx.Rollback()
}
