package postgresdriver

import (
	"context"

	"github.com/pokt-foundation/transaction-db/types"
)

func (d *PostgresDriver) WriteError(ctx context.Context, poktErr types.Error) error {
	return d.InsertError(ctx, InsertErrorParams{
		ErrorCode:        poktErr.ErrorCode,
		ErrorName:        poktErr.ErrorName,
		ErrorDescription: poktErr.ErrorDescription,
	})
}
