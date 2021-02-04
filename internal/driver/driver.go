package driver

import (
	"database/sql"
	"github.com/muktiarafi/myriadcode-backend/internal/configs"
	"os"
	"path/filepath"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

const (
	maxOpenDBConn = 10
	maxIdleDBCnn  = 5
	maxDBLifetime = 5 * time.Minute
)

// ConnectSQL creates database pool for Postgres
func ConnectSQL(dsn string, app *configs.AppConfig) (*DB, error) {
	d, err := newDatabase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDBConn)
	d.SetMaxIdleConns(maxIdleDBCnn)
	d.SetConnMaxLifetime(maxDBLifetime)

	dbConn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	if app.WithMigration {
		pwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		migrationFilePath := filepath.Join(pwd, "database", "migrations")

		app.InfoLog.Println("Running Migration")
		if err := Migration(migrationFilePath, d); err != nil {
			return nil, err
		}
	}

	return dbConn, nil
}

// newDatabase creates a new database for the application
func newDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	if err := testDB(db); err != nil {
		return nil, err
	}

	return db, nil
}

// testDB tries to ping the database
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		return err
	}

	return nil
}
