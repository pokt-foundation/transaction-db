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
	relay_id,
	pokt_chain_id,
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
	pokt_tx_id,
	created_at,
	updated_at
  )
  SELECT * FROM unnest(
	$1::varchar[],
	$2::varchar[],
	$3::varchar[],
	$4::varchar[],
	$5::varchar[],
	$6::varchar[],
	$7::varchar[],
	$8::varchar[],
	$9::date[],
	$10::date[],
	$11::boolean[],
	$12::integer[],
	$13::varchar[],
	$14::varchar[],
	$15::error_sources_enum[],
	$16::varchar[],
	$17::integer[],
	$18::varchar[],
	$19::integer[],
	$20::integer[],
	$21::integer[],
	$22::boolean[],
	$23::integer[],
	$24::boolean[],
	$25::boolean[],
	$26::varchar[],
	$27::varchar[],
	$28::date[],
	$29::date[]
  ) AS t(
	relay_id,
	pokt_chain_id,
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
	pokt_tx_id,
	created_at,
	updated_at
  )`

const chainMethodIDSeparator = ","

func (d *PostgresDriver) WriteRelay(ctx context.Context, relay types.Relay) error {
	now := time.Now()

	return d.InsertRelay(ctx, InsertRelayParams{
		RelayID:                  relay.RelayID,
		PoktChainID:              relay.PoktChainID,
		EndpointID:               relay.EndpointID,
		SessionKey:               relay.SessionKey,
		RelaySourceUrl:           relay.RelaySourceURL,
		PoktNodeAddress:          relay.PoktNodeAddress,
		PoktNodeDomain:           relay.PoktNodeDomain,
		PoktNodePublicKey:        relay.PoktNodePublicKey,
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorCode:                newSQLNullInt32(int32(relay.ErrorCode)),
		ErrorName:                newSQLNullString(relay.ErrorName),
		ErrorMessage:             newSQLNullString(relay.ErrorMessage),
		ErrorType:                newSQLNullString(relay.ErrorType),
		ErrorSource:              newSQLNullErrorSource(ErrorSourcesEnum(relay.ErrorSource)),
		RelayRoundtripTime:       int32(relay.RelayRoundtripTime),
		RelayChainMethodIds:      strings.Join(relay.RelayChainMethodIDs, chainMethodIDSeparator),
		RelayDataSize:            int32(relay.RelayDataSize),
		RelayPortalTripTime:      int32(relay.RelayPortalTripTime),
		RelayNodeTripTime:        int32(relay.RelayNodeTripTime),
		RelayUrlIsPublicEndpoint: relay.RelayURLIsPublicEndpoint,
		PortalOriginRegionID:     int32(relay.PortalOriginRegionID),
		IsAltruistRelay:          relay.IsAltruistRelay,
		RequestID:                relay.RequestID,
		PoktTxID:                 relay.PoktTxID,
		IsUserRelay:              relay.IsUserRelay,
		CreatedAt:                now,
		UpdatedAt:                now,
	})
}

func (d *PostgresDriver) WriteRelays(ctx context.Context, relays []types.Relay) error {
	now := time.Now()

	var (
		relayIDs                  []string
		poktChainIDs              []string
		endpointIDs               []string
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
		errorTypes                []sql.NullString
		errorSources              []NullErrorSourcesEnum
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
		poktTxIDs                 []string
		createdTimes              []time.Time
		updatedTimes              []time.Time
	)

	for _, relay := range relays {
		relayIDs = append(relayIDs, relay.RelayID)
		poktChainIDs = append(poktChainIDs, relay.PoktChainID)
		endpointIDs = append(endpointIDs, relay.EndpointID)
		sessionKeys = append(sessionKeys, relay.SessionKey)
		relaySourceURLs = append(relaySourceURLs, relay.RelaySourceURL)
		poktNodeAddresses = append(poktNodeAddresses, relay.PoktNodeAddress)
		poktNodeDomains = append(poktNodeDomains, relay.PoktNodeDomain)
		poktNodePublicKeys = append(poktNodePublicKeys, relay.PoktNodePublicKey)
		relayStartDatetimes = append(relayStartDatetimes, relay.RelayStartDatetime)
		relayReturnDatetimes = append(relayReturnDatetimes, relay.RelayReturnDatetime)
		isErrors = append(isErrors, relay.IsError)
		errorCodes = append(errorCodes, newSQLNullInt32(int32(relay.ErrorCode)))
		errorNames = append(errorNames, newSQLNullString(relay.ErrorName))
		errorMessages = append(errorMessages, newSQLNullString(relay.ErrorMessage))
		errorTypes = append(errorTypes, newSQLNullString(relay.ErrorType))
		errorSources = append(errorSources, newSQLNullErrorSource(ErrorSourcesEnum(relay.ErrorSource)))
		relayRoundtripTimes = append(relayRoundtripTimes, int32(relay.RelayRoundtripTime))
		relayChainMethodIDs = append(relayChainMethodIDs, strings.Join(relay.RelayChainMethodIDs, chainMethodIDSeparator))
		relayDataSizes = append(relayDataSizes, int32(relay.RelayDataSize))
		relayPortalTripTimes = append(relayPortalTripTimes, int32(relay.RelayPortalTripTime))
		relayNodeTripTimes = append(relayNodeTripTimes, int32(relay.RelayNodeTripTime))
		relayURLIsPublicEndpoints = append(relayURLIsPublicEndpoints, relay.RelayURLIsPublicEndpoint)
		portalOriginRegionIDs = append(portalOriginRegionIDs, int32(relay.PortalOriginRegionID))
		isAltruistRelays = append(isAltruistRelays, relay.IsAltruistRelay)
		isUserRelays = append(isUserRelays, relay.IsUserRelay)
		requestIDs = append(requestIDs, relay.RequestID)
		poktTxIDs = append(poktTxIDs, relay.PoktTxID)
		createdTimes = append(createdTimes, now)
		updatedTimes = append(updatedTimes, now)
	}

	_, err := d.db.Exec(insertRelays, pq.StringArray(relayIDs),
		pq.StringArray(poktChainIDs),
		pq.StringArray(endpointIDs),
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
		pq.StringArray(poktTxIDs),
		pq.Array(createdTimes),
		pq.Array(updatedTimes))
	if err != nil {
		return err
	}

	return nil
}

func (d *PostgresDriver) ReadRelay(ctx context.Context, relayID string) (types.Relay, error) {
	relay, err := d.SelectRelay(ctx, relayID)
	if err != nil {
		return types.Relay{}, err
	}

	return types.Relay{
		RelayID:                  relay.RelayID,
		PoktChainID:              relay.PoktChainID,
		EndpointID:               relay.EndpointID,
		SessionKey:               relay.SessionKey,
		RelaySourceURL:           relay.RelaySourceUrl,
		PoktNodeAddress:          relay.PoktNodeAddress,
		PoktNodeDomain:           relay.PoktNodeDomain,
		PoktNodePublicKey:        relay.PoktNodePublicKey,
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorCode:                int(relay.ErrorCode.Int32),
		ErrorName:                relay.ErrorName.String,
		ErrorMessage:             relay.ErrorMessage.String,
		ErrorType:                relay.ErrorType.String,
		ErrorSource:              types.ErrorSource(relay.ErrorSource.ErrorSourcesEnum),
		RelayRoundtripTime:       int(relay.RelayRoundtripTime),
		RelayChainMethodIDs:      strings.Split(relay.RelayChainMethodIds, ","),
		RelayDataSize:            int(relay.RelayDataSize),
		RelayPortalTripTime:      int(relay.RelayPortalTripTime),
		RelayNodeTripTime:        int(relay.RelayNodeTripTime),
		RelayURLIsPublicEndpoint: relay.RelayUrlIsPublicEndpoint,
		PortalOriginRegionID:     int(relay.PortalOriginRegionID),
		IsAltruistRelay:          relay.IsAltruistRelay,
		RequestID:                relay.RequestID,
		IsUserRelay:              relay.IsUserRelay,
		PoktTxID:                 relay.PoktTxID,
		CreatedAt:                relay.CreatedAt,
		UpdatedAt:                relay.UpdatedAt,
		Session: types.PocketSession{
			SessionKey:            relay.SessionKey,
			SessionHeight:         int(relay.SessionHeight),
			ProtocolApplicationID: relay.ProtocolApplicationID,
			ProtocolPublicKey:     relay.ProtocolPublicKey,
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
