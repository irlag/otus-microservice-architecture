package db

import (
	"database/sql"

	"otus-microservice-architecture/app/config"
)

// Store defines all functions to execute db queries and transactions
type Store interface {
	Open(config *config.DBConfig) error
	Close() error
	Querier
	Transactional
	WithTx(tx *sql.Tx) *Queries
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore() Store {
	return &SQLStore{}
}

func (s *SQLStore) Open(config *config.DBConfig) error {
	db, err := sql.Open(config.Driver, config.GetDSN())
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxConns)

	db.SetConnMaxIdleTime(config.MaxIdleConnTime)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)

	if err = db.Ping(); err != nil {
		return err
	}

	s.db = db
	s.Queries = New(db)

	return nil
}

func (s *SQLStore) Close() error {
	if err := s.db.Close(); err != nil {
		return err
	}

	return nil
}

func (s *SQLStore) Begin() (*sql.Tx, error) {
	return s.db.Begin()
}
