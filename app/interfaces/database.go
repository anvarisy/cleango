package interfaces

import "database/sql"

type DB interface {
	GetConnMySQL() (*sql.DB, error)
}
