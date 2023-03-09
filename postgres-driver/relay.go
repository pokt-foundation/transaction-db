package postgresdriver

import (
	"context"
	"time"

	"github.com/pokt-foundation/transaction-db/types"
)

func (d *PostgresDriver) WriteRelay(ctx context.Context, relay types.Relay) error {
	now := time.Now()

	return d.InsertRelay(ctx, InsertRelayParams{
		ChainID:                  relay.ChainID,
		EndpointID:               relay.EndpointID,
		PocketSessionID:          relay.PocketSessionID,
		PoktNodeAddress:          relay.PoktNodeAddress,
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorID:                  relay.ErrorID,
		RelayRoundtripTime:       relay.RelayRoundtripTime,
		RelayChainMethodID:       relay.RelayChainMethodID,
		RelayDataSize:            relay.RelayDataSize,
		RelayPortalTripTime:      relay.RelayPortalTripTime,
		RelayNodeTripTime:        relay.RelayNodeTripTime,
		RelayUrlIsPublicEndpoint: relay.RelayURLIsPublicEndpoint,
		PortalOriginRegionID:     relay.PortalOriginRegionID,
		IsAltruistRelay:          relay.IsAltruistRelay,
		CreatedAt:                now,
		UpdatedAt:                now,
	})
}

func (d *PostgresDriver) ReadRelay(ctx context.Context, relayID int) (types.Relay, error) {
	relay, err := d.SelectRelay(ctx, int64(relayID))
	if err != nil {
		return types.Relay{}, err
	}

	return types.Relay{
		RelayID:                  relay.RelayID,
		ChainID:                  relay.ChainID,
		EndpointID:               relay.EndpointID,
		PocketSessionID:          relay.PocketSessionID,
		PoktNodeAddress:          relay.PoktNodeAddress,
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorID:                  relay.ErrorID,
		RelayRoundtripTime:       relay.RelayRoundtripTime,
		RelayChainMethodID:       relay.RelayChainMethodID,
		RelayDataSize:            relay.RelayDataSize,
		RelayPortalTripTime:      relay.RelayPortalTripTime,
		RelayNodeTripTime:        relay.RelayNodeTripTime,
		RelayURLIsPublicEndpoint: relay.RelayUrlIsPublicEndpoint,
		PortalOriginRegionID:     relay.PortalOriginRegionID,
		IsAltruistRelay:          relay.IsAltruistRelay,
		CreatedAt:                relay.CreatedAt,
		UpdatedAt:                relay.UpdatedAt,
		Session: types.PocketSession{
			SessionKey:            relay.SessionKey,
			SessionHeight:         relay.SessionHeight,
			ProtocolApplicationID: relay.ProtocolApplicationID,
			CreatedAt:             relay.CreatedAt_2,
			UpdatedAt:             relay.UpdatedAt_2,
		},
		Region: types.PortalRegion{
			PortalRegionName: relay.PortalRegionName,
			CreatedAt:        relay.CreatedAt_3,
			UpdatedAt:        relay.UpdatedAt_3,
		},
		Error: types.Error{
			ErrorCode:        relay.ErrorCode,
			ErrorName:        relay.ErrorName,
			ErrorDescription: relay.ErrorDescription,
			CreatedAt:        relay.CreatedAt_4,
			UpdatedAt:        relay.UpdatedAt_4,
		},
	}, nil
}
