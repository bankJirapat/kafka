package database_function

import "kafka/pkg/common/db"

type DatabaseFunction struct {
	db *db.DB
}

func UseDatabase(db *db.DB) *DatabaseFunction {
	return &DatabaseFunction{
		db: db,
	}
}
