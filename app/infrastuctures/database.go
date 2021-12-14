package infrastuctures

import (
	"cleango/app/interfaces"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct{}

// NewDB creates a DB struct.
func NewDB() interfaces.DB {
	return &DB{}
}
func (d *DB) GetConnMySQL() (*sql.DB, error) {
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	conn, err := sql.Open("postgres", "postgresql://postgres:koalapanda@localhost:5432/sekolahimpian?sslmode=disable")
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Second * 10)

	return conn, nil
}
