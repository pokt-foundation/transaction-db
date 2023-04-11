package postgresdriver

import (
	"context"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

func (d *PostgresDriver) WriteSession(ctx context.Context, session types.PocketSession) error {
	now := time.Now()

	return d.InsertPocketSession(ctx, InsertPocketSessionParams{
		SessionKey:            session.SessionKey,
		SessionHeight:         int32(session.SessionHeight),
		ProtocolApplicationID: session.ProtocolApplicationID,
		ProtocolPublicKey:     session.ProtocolPublicKey,
		CreatedAt:             now,
		UpdatedAt:             now,
	})
}
