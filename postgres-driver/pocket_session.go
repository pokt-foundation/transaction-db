package postgresdriver

import (
	"context"

	"github.com/pokt-foundation/transaction-db/types"
)

func (d *PostgresDriver) WriteSession(ctx context.Context, session types.PocketSession) error {
	return d.InsertPocketSession(ctx, InsertPocketSessionParams{
		SessionKey:            session.SessionKey,
		SessionHeight:         session.SessionHeight,
		ProtocolApplicationID: session.ProtocolApplicationID,
	})
}
