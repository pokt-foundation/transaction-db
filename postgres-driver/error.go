package postgresdriver

import (
	"context"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

func (d *PostgresDriver) WriteError(ctx context.Context, poktErr types.Error) error {
	now := time.Now()

	return d.InsertError(ctx, InsertErrorParams{
		ErrorCode:        poktErr.ErrorCode,
		ErrorName:        poktErr.ErrorName,
		ErrorDescription: poktErr.ErrorDescription,
		ErrorType:        ErrorTypesEnum(poktErr.ErrorType),
		CreatedAt:        now,
		UpdatedAt:        now,
	})
}
