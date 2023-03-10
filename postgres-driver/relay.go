package postgresdriver

import (
	"context"
	"time"

	"github.com/lib/pq"
	"github.com/pokt-foundation/transaction-db/types"
)

const insertRelays = `INSERT INTO relay (
	chain_id,
	endpoint_id,
	pocket_session_id,
	pokt_node_address,
	relay_start_datetime,
	relay_return_datetime,
	is_error,
	error_id,
	relay_roundtrip_time,
	relay_chain_method_id,
	relay_data_size,
	relay_portal_trip_time,
	relay_node_trip_time,
	relay_url_is_public_endpoint,
	portal_origin_region_id,
	is_altruist_relay,
	created_at,
	updated_at
  )
  SELECT * FROM unnest(
	$1::integer[],
	$2::integer[],
	$3::bigint[],
	$4::varchar[],
	$5::date[],
	$6::date[],
	$7::boolean[],
	$8::bigint[],
	$9::integer[],
	$10::integer[],
	$11::integer[],
	$12::integer[],
	$13::integer[],
	$14::boolean[],
	$15::integer[],
	$16::boolean[],
	$17::date[],
	$18::date[]
  ) AS t(
	chain_id,
	endpoint_id,
	pocket_session_id,
	pokt_node_address,
	relay_start_datetime,
	relay_return_datetime,
	is_error,
	error_id,
	relay_roundtrip_time,
	relay_chain_method_id,
	relay_data_size,
	relay_portal_trip_time,
	relay_node_trip_time,
	relay_url_is_public_endpoint,
	portal_origin_region_id,
	is_altruist_relay,
	created_at,
	updated_at
  )`

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

func (d *PostgresDriver) WriteRelays(ctx context.Context, relays []types.Relay) error {
	now := time.Now()

	var (
		chainIDs                  []int32
		endpointIDs               []int32
		pocketSessionIDs          []int64
		poktNodeAddresses         []string
		relayStartDatetimes       []time.Time
		relayReturnDatetimes      []time.Time
		isErrors                  []bool
		errorIDs                  []int64
		relayRoundtripTimes       []int32
		relayChainMethodIDs       []int32
		relayDataSizes            []int32
		relayPortalTripTimes      []int32
		relayNodeTripTimes        []int32
		relayURLIsPublicEndpoints []bool
		portalOriginRegionIDs     []int32
		isAltruistRelays          []bool
		createdTimes              []time.Time
		updatedTimes              []time.Time
	)

	for _, relay := range relays {
		chainIDs = append(chainIDs, relay.ChainID)
		endpointIDs = append(endpointIDs, relay.EndpointID)
		pocketSessionIDs = append(pocketSessionIDs, relay.PocketSessionID)
		poktNodeAddresses = append(poktNodeAddresses, relay.PoktNodeAddress)
		relayStartDatetimes = append(relayStartDatetimes, relay.RelayStartDatetime)
		relayReturnDatetimes = append(relayReturnDatetimes, relay.RelayReturnDatetime)
		isErrors = append(isErrors, relay.IsError)
		errorIDs = append(errorIDs, relay.ErrorID)
		relayRoundtripTimes = append(relayRoundtripTimes, relay.RelayRoundtripTime)
		relayChainMethodIDs = append(relayChainMethodIDs, relay.RelayChainMethodID)
		relayDataSizes = append(relayDataSizes, relay.RelayDataSize)
		relayPortalTripTimes = append(relayPortalTripTimes, relay.RelayPortalTripTime)
		relayNodeTripTimes = append(relayNodeTripTimes, relay.RelayNodeTripTime)
		relayURLIsPublicEndpoints = append(relayURLIsPublicEndpoints, relay.RelayURLIsPublicEndpoint)
		portalOriginRegionIDs = append(portalOriginRegionIDs, relay.PortalOriginRegionID)
		isAltruistRelays = append(isAltruistRelays, relay.IsAltruistRelay)
		createdTimes = append(createdTimes, now)
		updatedTimes = append(updatedTimes, now)
	}

	_, err := d.db.Exec(insertRelays, pq.Int32Array(chainIDs),
		pq.Int32Array(endpointIDs),
		pq.Int64Array(pocketSessionIDs),
		pq.StringArray(poktNodeAddresses),
		pq.Array(relayStartDatetimes),
		pq.Array(relayReturnDatetimes),
		pq.BoolArray(isErrors),
		pq.Int64Array(errorIDs),
		pq.Int32Array(relayRoundtripTimes),
		pq.Int32Array(relayChainMethodIDs),
		pq.Int32Array(relayDataSizes),
		pq.Int32Array(relayPortalTripTimes),
		pq.Int32Array(relayNodeTripTimes),
		pq.BoolArray(relayURLIsPublicEndpoints),
		pq.Int32Array(portalOriginRegionIDs),
		pq.BoolArray(isAltruistRelays),
		pq.Array(createdTimes),
		pq.Array(updatedTimes))
	if err != nil {
		return err
	}

	return nil
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
			ErrorType:        types.ErrorType(relay.ErrorType),
			CreatedAt:        relay.CreatedAt_4,
			UpdatedAt:        relay.UpdatedAt_4,
		},
	}, nil
}
