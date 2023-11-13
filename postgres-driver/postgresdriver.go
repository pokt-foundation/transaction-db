package postgresdriver

import (
	"context"
	"fmt"
	"time"

	// PQ import is required

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// The PostgresDriver struct satisfies the Driver interface which defines all database driver methods
type PostgresDriver struct {
	*Queries
	db *pgxpool.Pool
}

/* ---------- Postgres Connection Funcs ---------- */

/* <--------- PGX Pool Connection ---------> */

/*
NewPostgresDriver
- Creates a pool of connections to a PostgreSQL database using the provided connection string.
- Parses the connection string into a pgx pool configuration object.
- For each acquired connection from the pool, custom enum types are registered.
- Returns the established connection pool.
- This function is ideal for creating multiple reusable connections to a PostgreSQL database, particularly useful for handling multiple concurrent database operations.
*/
func NewPostgresDriver(connectionString string) (*PostgresDriver, func() error, error) {
	config, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, nil, err
	}

	pool, err := createAndConfigurePool(config)
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() error {
		pool.Close()
		return nil
	}

	driver := &PostgresDriver{
		Queries: New(pool),
		db:      pool,
	}

	return driver, cleanup, nil
}

// Configures the connection pool with custom enum types.
func createAndConfigurePool(config *pgxpool.Config) (*pgxpool.Pool, error) {
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.NewWithConfig: %v", err)
	}

	// Collect the custom data types once, store them in memory, and register them for every future connection.
	customTypes, err := getCustomDataTypes(context.Background(), pool)
	if err != nil {
		return nil, err
	}
	config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		for _, t := range customTypes {
			conn.TypeMap().RegisterType(t)
		}
		return nil
	}

	// Immediately close the old pool and open a new one with the new config.
	pool.Close()
	pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

// Any custom DB types made with CREATE TYPE need to be registered with pgx.
// https://github.com/kyleconroy/sqlc/issues/2116
// https://stackoverflow.com/questions/75658429/need-to-update-psql-row-of-a-composite-type-in-golang-with-jack-pgx
// https://pkg.go.dev/github.com/jackc/pgx/v5/pgtype
func getCustomDataTypes(ctx context.Context, pool *pgxpool.Pool) ([]*pgtype.Type, error) {
	// Get a single connection just to load type information.
	conn, err := pool.Acquire(ctx)
	defer conn.Release()
	if err != nil {
		return nil, err
	}

	dataTypeNames := []string{
		"error_sources_enum",
		"_error_sources_enum",
	}

	var typesToRegister []*pgtype.Type
	for _, typeName := range dataTypeNames {
		dataType, err := conn.Conn().LoadType(ctx, typeName)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to load type %s: %v", typeName, err)
		}
		// You need to register only for this connection too, otherwise the array type will look for the register element type.
		conn.Conn().TypeMap().RegisterType(dataType)
		typesToRegister = append(typesToRegister, dataType)
	}
	return typesToRegister, nil
}

func newText(value string) pgtype.Text {
	if value == "" {
		return pgtype.Text{}
	}

	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func newInt4(value int32, allowZero bool) pgtype.Int4 {
	if !allowZero && value == 0 {
		return pgtype.Int4{}
	}

	return pgtype.Int4{
		Int32: value,
		Valid: true,
	}
}

func newTimestamp(value time.Time) pgtype.Timestamp {
	if value.IsZero() {
		return pgtype.Timestamp{}
	}

	return pgtype.Timestamp{
		Time:  value,
		Valid: true,
	}
}

func newNullErrorSourcesEnum(e ErrorSourcesEnum) NullErrorSourcesEnum {
	if e == "" {
		return NullErrorSourcesEnum{}
	}

	return NullErrorSourcesEnum{
		ErrorSourcesEnum: e,
		Valid:            true,
	}
}
