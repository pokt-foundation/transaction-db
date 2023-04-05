package postgresdriver

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/lib/pq"
	"github.com/pokt-foundation/transaction-db/types"
)

const insertRelays = `INSERT INTO relay (
	chain_id,
	endpoint_id,
	session_key,
	relay_source_url,
	pokt_node_address,
	pokt_node_domain,
	pokt_node_public_key,
	relay_start_datetime,
	relay_return_datetime,
	is_error,
	error_code,
	error_name,
	error_message,
	error_source,
	error_type,
	relay_roundtrip_time,
	relay_chain_method_ids,
	relay_data_size,
	relay_portal_trip_time,
	relay_node_trip_time,
	relay_url_is_public_endpoint,
	portal_origin_region_id,
	is_altruist_relay,
	is_user_relay,
	request_id,
	created_at,
	updated_at
  )
  SELECT * FROM unnest(
	$1::integer[],
	$2::integer[],
	$3::varchar[],
	$4::varchar[],
	$5::varchar[],
	$6::varchar[],
	$7::varchar[],
	$8::date[],
	$9::date[],
	$10::boolean[],
	$11::integer[],
	$12::varchar[],
	$13::varchar[],
	$14::varchar[],
	$15::error_types_enum[],
	$16::integer[],
	$17::varchar[],
	$18::integer[],
	$19::integer[],
	$20::integer[],
	$21::boolean[],
	$22::integer[],
	$23::boolean[],
	$24::boolean[],
	$25::varchar[],
	$26::date[],
	$27::date[]
  ) AS t(
	chain_id,
	endpoint_id,
	session_key,
	relay_source_url,
	pokt_node_address,
	pokt_node_domain,
	pokt_node_public_key,
	relay_start_datetime,
	relay_return_datetime,
	is_error,
	error_code,
	error_name,
	error_message,
	error_source,
	error_type,
	relay_roundtrip_time,
	relay_chain_method_ids,
	relay_data_size,
	relay_portal_trip_time,
	relay_node_trip_time,
	relay_url_is_public_endpoint,
	portal_origin_region_id,
	is_altruist_relay,
	is_user_relay,
	request_id,
	created_at,
	updated_at
  )`

const chainMethodIDSeparator = ","

func (d *PostgresDriver) WriteRelay(ctx context.Context, relay types.Relay) error {
	now := time.Now()

	return d.InsertRelay(ctx, InsertRelayParams{
		ChainID:                  relay.ChainID,
		EndpointID:               relay.EndpointID,
		SessionKey:               relay.SessionKey,
		RelaySourceUrl:           relay.RelaySourceURL,
		PoktNodeAddress:          relay.PoktNodeAddress,
		PoktNodeDomain:           relay.PoktNodeDomain,
		PoktNodePublicKey:        relay.PoktNodePublicKey,
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorCode:                newSQLNullInt32(relay.ErrorCode),
		ErrorName:                newSQLNullString(relay.ErrorName),
		ErrorMessage:             newSQLNullString(relay.ErrorMessage),
		ErrorType:                newSQLNullErrorType(ErrorTypesEnum(relay.ErrorType)),
		ErrorSource:              newSQLNullString(relay.ErrorSource),
		RelayRoundtripTime:       relay.RelayRoundtripTime,
		RelayChainMethodIds:      strings.Join(relay.RelayChainMethodIDs, chainMethodIDSeparator),
		RelayDataSize:            relay.RelayDataSize,
		RelayPortalTripTime:      relay.RelayPortalTripTime,
		RelayNodeTripTime:        relay.RelayNodeTripTime,
		RelayUrlIsPublicEndpoint: relay.RelayURLIsPublicEndpoint,
		PortalOriginRegionID:     relay.PortalOriginRegionID,
		IsAltruistRelay:          relay.IsAltruistRelay,
		RequestID:                relay.RequestID,
		IsUserRelay:              relay.IsUserRelay,
		CreatedAt:                now,
		UpdatedAt:                now,
	})
}

func (d *PostgresDriver) WriteRelays(ctx context.Context, relays []types.Relay) error {
	now := time.Now()

	var (
		chainIDs                  []int32
		endpointIDs               []int32
		sessionKeys               []string
		relaySourceURLs           []string
		poktNodeAddresses         []string
		poktNodeDomains           []string
		poktNodePublicKeys        []string
		relayStartDatetimes       []time.Time
		relayReturnDatetimes      []time.Time
		isErrors                  []bool
		errorCodes                []sql.NullInt32
		errorNames                []sql.NullString
		errorMessages             []sql.NullString
		errorTypes                []NullErrorTypesEnum
		errorSources              []sql.NullString
		relayRoundtripTimes       []int32
		relayChainMethodIDs       []string
		relayDataSizes            []int32
		relayPortalTripTimes      []int32
		relayNodeTripTimes        []int32
		relayURLIsPublicEndpoints []bool
		portalOriginRegionIDs     []int32
		isAltruistRelays          []bool
		isUserRelays              []bool
		requestIDs                []string
		createdTimes              []time.Time
		updatedTimes              []time.Time
	)

	for _, relay := range relays {
		chainIDs = append(chainIDs, relay.ChainID)
		endpointIDs = append(endpointIDs, relay.EndpointID)
		sessionKeys = append(sessionKeys, relay.SessionKey)
		relaySourceURLs = append(relaySourceURLs, relay.RelaySourceURL)
		poktNodeAddresses = append(poktNodeAddresses, relay.PoktNodeAddress)
		poktNodeDomains = append(poktNodeDomains, relay.PoktNodeDomain)
		poktNodePublicKeys = append(poktNodePublicKeys, relay.PoktNodePublicKey)
		relayStartDatetimes = append(relayStartDatetimes, relay.RelayStartDatetime)
		relayReturnDatetimes = append(relayReturnDatetimes, relay.RelayReturnDatetime)
		isErrors = append(isErrors, relay.IsError)
		errorCodes = append(errorCodes, newSQLNullInt32(relay.ErrorCode))
		errorNames = append(errorNames, newSQLNullString(relay.ErrorName))
		errorMessages = append(errorMessages, newSQLNullString(relay.ErrorMessage))
		errorTypes = append(errorTypes, newSQLNullErrorType(ErrorTypesEnum(relay.ErrorType)))
		errorSources = append(errorSources, newSQLNullString(relay.ErrorSource))
		relayRoundtripTimes = append(relayRoundtripTimes, relay.RelayRoundtripTime)
		relayChainMethodIDs = append(relayChainMethodIDs, strings.Join(relay.RelayChainMethodIDs, chainMethodIDSeparator))
		relayDataSizes = append(relayDataSizes, relay.RelayDataSize)
		relayPortalTripTimes = append(relayPortalTripTimes, relay.RelayPortalTripTime)
		relayNodeTripTimes = append(relayNodeTripTimes, relay.RelayNodeTripTime)
		relayURLIsPublicEndpoints = append(relayURLIsPublicEndpoints, relay.RelayURLIsPublicEndpoint)
		portalOriginRegionIDs = append(portalOriginRegionIDs, relay.PortalOriginRegionID)
		isAltruistRelays = append(isAltruistRelays, relay.IsAltruistRelay)
		isUserRelays = append(isUserRelays, relay.IsUserRelay)
		requestIDs = append(requestIDs, relay.RequestID)
		createdTimes = append(createdTimes, now)
		updatedTimes = append(updatedTimes, now)
	}

	_, err := d.db.Exec(insertRelays, pq.Int32Array(chainIDs),
		pq.Int32Array(endpointIDs),
		pq.StringArray(sessionKeys),
		pq.StringArray(relaySourceURLs),
		pq.StringArray(poktNodeAddresses),
		pq.StringArray(poktNodeDomains),
		pq.StringArray(poktNodePublicKeys),
		pq.Array(relayStartDatetimes),
		pq.Array(relayReturnDatetimes),
		pq.BoolArray(isErrors),
		pq.Array(errorCodes),
		pq.Array(errorNames),
		pq.Array(errorMessages),
		pq.Array(errorSources),
		pq.Array(errorTypes),
		pq.Int32Array(relayRoundtripTimes),
		pq.StringArray(relayChainMethodIDs),
		pq.Int32Array(relayDataSizes),
		pq.Int32Array(relayPortalTripTimes),
		pq.Int32Array(relayNodeTripTimes),
		pq.BoolArray(relayURLIsPublicEndpoints),
		pq.Int32Array(portalOriginRegionIDs),
		pq.BoolArray(isAltruistRelays),
		pq.BoolArray(isUserRelays),
		pq.StringArray(requestIDs),
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
		SessionKey:               relay.SessionKey,
		RelaySourceURL:           relay.RelaySourceUrl,
		PoktNodeAddress:          relay.PoktNodeAddress,
		PoktNodeDomain:           relay.PoktNodeDomain,
		PoktNodePublicKey:        relay.PoktNodePublicKey,
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorCode:                relay.ErrorCode.Int32,
		ErrorName:                relay.ErrorName.String,
		ErrorMessage:             relay.ErrorMessage.String,
		ErrorType:                types.ErrorType(relay.ErrorType.ErrorTypesEnum),
		ErrorSource:              relay.ErrorSource.String,
		RelayRoundtripTime:       relay.RelayRoundtripTime,
		RelayChainMethodIDs:      strings.Split(relay.RelayChainMethodIds, ","),
		RelayDataSize:            relay.RelayDataSize,
		RelayPortalTripTime:      relay.RelayPortalTripTime,
		RelayNodeTripTime:        relay.RelayNodeTripTime,
		RelayURLIsPublicEndpoint: relay.RelayUrlIsPublicEndpoint,
		PortalOriginRegionID:     relay.PortalOriginRegionID,
		IsAltruistRelay:          relay.IsAltruistRelay,
		RequestID:                relay.RequestID,
		IsUserRelay:              relay.IsUserRelay,
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
	}, nil
}
