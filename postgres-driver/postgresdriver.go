package postgresdriver

import (
	"database/sql"

	// PQ import is required
	_ "github.com/lib/pq"
)

// The PostgresDriver struct satisfies the Driver interface which defines all database driver methods
type PostgresDriver struct {
	*Queries
	db *sql.DB
}

/* NewPostgresDriver returns PostgresDriver instance from Postgres connection string */
func NewPostgresDriver(connectionString string) (*PostgresDriver, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	driver := &PostgresDriver{
		Queries: New(db),
		db:      db,
	}

	return driver, nil
}

/* NewPostgresDriverFromDBInstance returns PostgresDriver instance from sdl.DB instance */
// mostly used for mocking tests
func NewPostgresDriverFromDBInstance(db *sql.DB) *PostgresDriver {
	driver := &PostgresDriver{
		Queries: New(db),
	}

	return driver
}

func newSQLNullInt32(value int32) sql.NullInt32 {
	if value == 0 {
		return sql.NullInt32{}
	}

	return sql.NullInt32{
		Int32: value,
		Valid: true,
	}
}

func newSQLNullString(value string) sql.NullString {
	if value == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: value,
		Valid:  true,
	}
}

func newSQLNullErrorType(value ErrorTypesEnum) NullErrorTypesEnum {
	if value == "" {
		return NullErrorTypesEnum{}
	}

	return NullErrorTypesEnum{
		ErrorTypesEnum: value,
		Valid:          true,
	}
}
