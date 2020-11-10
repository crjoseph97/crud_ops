package backend

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //mysql driver
	"github.com/jmoiron/sqlx"
)

// ConnectDB initializes mysql/postgres DB
func ConnectDB(databaseType, datasource string, maxactive, maxidle int) (*sqlx.DB, error) {
	db := sqlx.MustOpen(databaseType, datasource)
	db.SetMaxOpenConns(maxactive)
	db.SetMaxIdleConns(maxidle)

	err := db.Ping()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to mysql: %s err: %s", datasource, err)
	}
	return db, nil
}
