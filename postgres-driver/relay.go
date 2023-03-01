package postgresdriver

import (
	"context"

	"github.com/pokt-foundation/transaction-db/types"
)

func (d *PostgresDriver) WriteRelay(ctx context.Context, relay types.Relay) error {
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
		Error: types.Error{
			ErrorCode:        relay.ErrorCode,
			ErrorName:        relay.ErrorName,
			ErrorDescription: relay.ErrorDescription,
		},
		Session: types.PocketSession{
			SessionKey:            relay.SessionKey,
			SessionHeight:         relay.SessionHeight,
			ProtocolApplicationID: relay.ProtocolApplicationID,
		},
		Region: types.PortalRegion{
			PortalRegionName: relay.PortalRegionName,
		},
	}, nil
}
