package db

import (
	"github.com/jmoiron/sqlx"
)

// connPool is a pull of all recorded connections
var connPool map[string]*Conn

// Conn represents single connection to DB
type Conn struct {
	DSN    string
	Driver string
	sqlx.DB
}

// Driver represents type of a database
type Driver string

// Values for driver but only MYSQL used yet
const (
	MYSQL  = "mysql"
	PGSQL  = "postgres"
	SQLITE = "sqlite3"
)

func init() {
	connPool = make(map[string]*Conn)
}

// Add connection to a pool of connections
func Add(connName, driverName, dataSourceName string) error {
	db, err := sqlx.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	connPool[connName] = &Conn{DB: *db, DSN: dataSourceName, Driver: driverName}
	return nil
}

// Use established connection from the pool.
func Use(name string) *Conn {
	return connPool[name]
}
