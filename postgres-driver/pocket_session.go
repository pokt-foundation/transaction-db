package postgresdriver

import (
	"context"
	"errors"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

const (
	errMessageDuplicateSessionKey = `duplicate key value violates unique constraint "pocket_session_session_key_key"`
)

var (
	ErrRepeatedSessionKey = errors.New("repeated session key")
)

func (d *PostgresDriver) WriteSession(ctx context.Context, session types.PocketSession) error {
	now := time.Now()

	err := d.InsertPocketSession(ctx, InsertPocketSessionParams{
		SessionKey:       session.SessionKey,
		SessionHeight:    int32(session.SessionHeight),
		PortalRegionName: session.PortalRegionName,
		CreatedAt:        now,
		UpdatedAt:        now,
	})
	if err != nil {
		if isSpecifiedPqError(errMessageDuplicateSessionKey, err) {
			return ErrRepeatedSessionKey
		}

		return err
	}

	return nil
}
