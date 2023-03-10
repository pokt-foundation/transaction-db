package postgresdriver

import (
	"context"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

func (d *PostgresDriver) WriteRegion(ctx context.Context, region types.PortalRegion) error {
	now := time.Now()

	return d.InsertPortalRegion(ctx, InsertPortalRegionParams{
		PortalRegionName: region.PortalRegionName,
		CreatedAt:        now,
		UpdatedAt:        now,
	})
}
