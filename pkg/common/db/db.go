package db

import (
	"errors"

	"github.com/uptrace/bun"
)

type DB struct {
	*bun.DB
}

func GetDatabase(cfg *Config) (*DB, error) {
	switch cfg.Driver {
	case string(POSTGRES):
		return NewPostgresDatabase(cfg), nil
	}

	return nil, errors.New("Driver not found")
}

func TestConnection(db *DB) error {
	rows, err := db.Query("SELECT 1+1")
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
