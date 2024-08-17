package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type Database struct {
	db *sql.DB
}

func InitDB(connStr string) (*Database, error) {
	var err error
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
