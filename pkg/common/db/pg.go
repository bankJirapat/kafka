package db

import (
	"bytes"
	"database/sql"
	"fmt"

	"github.com/oiime/logrusbun"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	dbConn *DB
)

func SQLConnectionStringBuilder(cfg *Config) string {
	return connectionStringBuilder(cfg)
}

func connectionStringBuilder(cfg *Config) string {
	var bufString bytes.Buffer
	bufString.WriteString(cfg.Driver)
	bufString.WriteString("://")
	bufString.WriteString(cfg.User)
	bufString.WriteString(":")
	bufString.WriteString(cfg.Password)
	bufString.WriteString("@")
	bufString.WriteString(cfg.Host)
	bufString.WriteString(":")
	bufString.WriteString(cfg.Port)
	bufString.WriteString("/")
	bufString.WriteString(cfg.Name)
	bufString.WriteString("?sslmode=")
	bufString.WriteString(cfg.SSLEnable)
	fmt.Printf("bufString : %+v\n", bufString.String())
	return bufString.String()
}

func connect(Path string) *bun.DB {

	// Open a PostgreSQL database.
	dsn := Path
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Create a Bun db on top of it.
	db := bun.NewDB(pgdb, pgdialect.New())

	return db
}

func NewPostgresDatabase(cfg *Config) *DB {

	dbConn = newPostgresDatabase(cfg)
	return dbConn
}

func newPostgresDatabase(cfg *Config) *DB {
	db := connect(connectionStringBuilder(cfg))
	fmt.Printf("db : %+v\n", db)
	// Print all queries to stdout.

	db.AddQueryHook(logrusbun.NewQueryHook(logrusbun.QueryHookOptions{Logger: logrus.New()}))

	_db := &DB{DB: db}
	if err := TestConnection(_db); err != nil {
		panic(err)
	}
	return _db
}
