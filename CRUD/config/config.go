package config

import (
	"database/sql"
	// import driver
	_ "github.com/go-sql-driver/mysql"
)

// GetMySQLDB is public
func GetMySQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "vahid"
	dbPass := "7731531361?vP"
	dbName := "vp_demo_1"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
