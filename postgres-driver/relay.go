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
	pokt_chain_id,
	endpoint_id,
	session_key,
	protocol_app_public_key,
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
	portal_region_name,
	is_altruist_relay,
	is_user_relay,
	request_id,
	pokt_tx_id,
	gigastake_app_id,
	created_at,
	updated_at,
	blocking_plugin
  )
  SELECT * FROM unnest(
	$1::char(4)[],
	$2::varchar[],
	$3::char(44)[],
	$4::char(64)[],
	$5::varchar[],
	$6::char(40)[],
	$7::varchar[],
	$8::char(64)[],
	$9::timestamp[],
	$10::timestamp[],
	$11::boolean[],
	$12::integer[],
	$13::varchar[],
	$14::varchar[],
	$15::error_sources_enum[],
	$16::varchar[],
	$17::float[],
	$18::varchar[],
	$19::integer[],
	$20::float[],
	$21::float[],
	$22::boolean[],
	$23::varchar[],
	$24::boolean[],
	$25::boolean[],
	$26::varchar[],
	$27::varchar[],
	$28::varchar[],
	$29::timestamp[],
	$30::timestamp[],
	$31::varchar[]
  ) AS t(
	pokt_chain_id,
	endpoint_id,
	session_key,
	protocol_app_public_key,
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
	portal_region_name,
	is_altruist_relay,
	is_user_relay,
	request_id,
	pokt_tx_id,
	gigastake_app_id,
	created_at,
	updated_at,
	blocking_plugin
  )`

const chainMethodIDSeparator = ","

func (d *PostgresDriver) WriteRelay(ctx context.Context, relay types.Relay) error {
	now := time.Now()

	return d.InsertRelay(ctx, InsertRelayParams{
		PoktChainID:              relay.PoktChainID,
		EndpointID:               relay.EndpointID,
		SessionKey:               relay.SessionKey,
		ProtocolAppPublicKey:     relay.ProtocolAppPublicKey,
		RelaySourceUrl:           newSQLNullString(relay.RelaySourceURL),
		PoktNodeAddress:          newSQLNullString(relay.PoktNodeAddress),
		PoktNodeDomain:           newSQLNullString(relay.PoktNodeDomain),
		PoktNodePublicKey:        newSQLNullString(relay.PoktNodePublicKey),
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorCode:                newSQLNullInt32(int32(relay.ErrorCode)),
		ErrorName:                newSQLNullString(relay.ErrorName),
		ErrorMessage:             newSQLNullString(relay.ErrorMessage),
		ErrorType:                newSQLNullString(relay.ErrorType),
		ErrorSource:              newSQLNullErrorSource(ErrorSourcesEnum(relay.ErrorSource)),
		RelayRoundtripTime:       relay.RelayRoundtripTime,
		RelayChainMethodIds:      strings.Join(relay.RelayChainMethodIDs, chainMethodIDSeparator),
		RelayDataSize:            int32(relay.RelayDataSize),
		RelayPortalTripTime:      relay.RelayPortalTripTime,
		RelayNodeTripTime:        relay.RelayNodeTripTime,
		RelayUrlIsPublicEndpoint: relay.RelayURLIsPublicEndpoint,
		PortalRegionName:         relay.PortalRegionName,
		IsAltruistRelay:          relay.IsAltruistRelay,
		RequestID:                relay.RequestID,
		PoktTxID:                 newSQLNullString(relay.PoktTxID),
		IsUserRelay:              relay.IsUserRelay,
		GigastakeAppID:           newSQLNullString(relay.GigastakeAppID),
		CreatedAt:                now,
		UpdatedAt:                now,
		BlockingPlugin:           newSQLNullString(relay.BlockingPlugin),
	})
}

func (d *PostgresDriver) WriteRelays(ctx context.Context, relays []*types.Relay) error {
	now := time.Now()

	var (
		poktChainIDs              []string
		endpointIDs               []string
		sessionKeys               []string
		protocolAppPublicKeys     []string
		relaySourceURLs           []sql.NullString
		poktNodeAddresses         []sql.NullString
		poktNodeDomains           []sql.NullString
		poktNodePublicKeys        []sql.NullString
		relayStartDatetimes       []time.Time
		relayReturnDatetimes      []time.Time
		isErrors                  []bool
		errorCodes                []sql.NullInt32
		errorNames                []sql.NullString
		errorMessages             []sql.NullString
		errorTypes                []sql.NullString
		errorSources              []NullErrorSourcesEnum
		relayRoundtripTimes       []float64
		relayChainMethodIDs       []string
		relayDataSizes            []int32
		relayPortalTripTimes      []float64
		relayNodeTripTimes        []float64
		relayURLIsPublicEndpoints []bool
		portalRegionNames         []string
		isAltruistRelays          []bool
		isUserRelays              []bool
		requestIDs                []string
		poktTxIDs                 []sql.NullString
		gigastakeAppIDs           []sql.NullString
		createdTimes              []time.Time
		updatedTimes              []time.Time
		blockingPlugins           []sql.NullString
	)

	for _, relay := range relays {
		poktChainIDs = append(poktChainIDs, relay.PoktChainID)
		endpointIDs = append(endpointIDs, relay.EndpointID)
		sessionKeys = append(sessionKeys, relay.SessionKey)
		protocolAppPublicKeys = append(protocolAppPublicKeys, relay.ProtocolAppPublicKey)
		relaySourceURLs = append(relaySourceURLs, newSQLNullString(relay.RelaySourceURL))
		poktNodeAddresses = append(poktNodeAddresses, newSQLNullString(relay.PoktNodeAddress))
		poktNodeDomains = append(poktNodeDomains, newSQLNullString(relay.PoktNodeDomain))
		poktNodePublicKeys = append(poktNodePublicKeys, newSQLNullString(relay.PoktNodePublicKey))
		relayStartDatetimes = append(relayStartDatetimes, relay.RelayStartDatetime)
		relayReturnDatetimes = append(relayReturnDatetimes, relay.RelayReturnDatetime)
		isErrors = append(isErrors, relay.IsError)
		errorCodes = append(errorCodes, newSQLNullInt32(int32(relay.ErrorCode)))
		errorNames = append(errorNames, newSQLNullString(relay.ErrorName))
		errorMessages = append(errorMessages, newSQLNullString(relay.ErrorMessage))
		errorTypes = append(errorTypes, newSQLNullString(relay.ErrorType))
		errorSources = append(errorSources, newSQLNullErrorSource(ErrorSourcesEnum(relay.ErrorSource)))
		relayRoundtripTimes = append(relayRoundtripTimes, relay.RelayRoundtripTime)
		relayChainMethodIDs = append(relayChainMethodIDs, strings.Join(relay.RelayChainMethodIDs, chainMethodIDSeparator))
		relayDataSizes = append(relayDataSizes, int32(relay.RelayDataSize))
		relayPortalTripTimes = append(relayPortalTripTimes, relay.RelayPortalTripTime)
		relayNodeTripTimes = append(relayNodeTripTimes, relay.RelayNodeTripTime)
		relayURLIsPublicEndpoints = append(relayURLIsPublicEndpoints, relay.RelayURLIsPublicEndpoint)
		portalRegionNames = append(portalRegionNames, relay.PortalRegionName)
		isAltruistRelays = append(isAltruistRelays, relay.IsAltruistRelay)
		isUserRelays = append(isUserRelays, relay.IsUserRelay)
		requestIDs = append(requestIDs, relay.RequestID)
		poktTxIDs = append(poktTxIDs, newSQLNullString(relay.PoktTxID))
		gigastakeAppIDs = append(gigastakeAppIDs, newSQLNullString(relay.GigastakeAppID))
		createdTimes = append(createdTimes, now)
		updatedTimes = append(updatedTimes, now)
		blockingPlugins = append(blockingPlugins, newSQLNullString(relay.BlockingPlugin))
	}

	_, err := d.db.Exec(insertRelays, pq.StringArray(poktChainIDs),
		pq.StringArray(endpointIDs),
		pq.StringArray(sessionKeys),
		pq.StringArray(protocolAppPublicKeys),
		pq.Array(relaySourceURLs),
		pq.Array(poktNodeAddresses),
		pq.Array(poktNodeDomains),
		pq.Array(poktNodePublicKeys),
		pq.Array(relayStartDatetimes),
		pq.Array(relayReturnDatetimes),
		pq.BoolArray(isErrors),
		pq.Array(errorCodes),
		pq.Array(errorNames),
		pq.Array(errorMessages),
		pq.Array(errorSources),
		pq.Array(errorTypes),
		pq.Float64Array(relayRoundtripTimes),
		pq.StringArray(relayChainMethodIDs),
		pq.Int32Array(relayDataSizes),
		pq.Float64Array(relayPortalTripTimes),
		pq.Float64Array(relayNodeTripTimes),
		pq.BoolArray(relayURLIsPublicEndpoints),
		pq.StringArray(portalRegionNames),
		pq.BoolArray(isAltruistRelays),
		pq.BoolArray(isUserRelays),
		pq.StringArray(requestIDs),
		pq.Array(poktTxIDs),
		pq.Array(gigastakeAppIDs),
		pq.Array(createdTimes),
		pq.Array(updatedTimes),
		pq.Array(blockingPlugins))
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
		RelayID:                  int(relay.ID),
		PoktChainID:              relay.PoktChainID,
		EndpointID:               relay.EndpointID,
		SessionKey:               relay.SessionKey,
		ProtocolAppPublicKey:     relay.ProtocolAppPublicKey,
		RelaySourceURL:           relay.RelaySourceUrl.String,
		PoktNodeAddress:          relay.PoktNodeAddress.String,
		PoktNodeDomain:           relay.PoktNodeDomain.String,
		PoktNodePublicKey:        relay.PoktNodePublicKey.String,
		RelayStartDatetime:       relay.RelayStartDatetime,
		RelayReturnDatetime:      relay.RelayReturnDatetime,
		IsError:                  relay.IsError,
		ErrorCode:                int(relay.ErrorCode.Int32),
		ErrorName:                relay.ErrorName.String,
		ErrorMessage:             relay.ErrorMessage.String,
		ErrorType:                relay.ErrorType.String,
		ErrorSource:              types.ErrorSource(relay.ErrorSource.ErrorSourcesEnum),
		RelayRoundtripTime:       relay.RelayRoundtripTime,
		RelayChainMethodIDs:      strings.Split(relay.RelayChainMethodIds, ","),
		RelayDataSize:            int(relay.RelayDataSize),
		RelayPortalTripTime:      relay.RelayPortalTripTime,
		RelayNodeTripTime:        relay.RelayNodeTripTime,
		RelayURLIsPublicEndpoint: relay.RelayUrlIsPublicEndpoint,
		PortalRegionName:         relay.PortalRegionName,
		IsAltruistRelay:          relay.IsAltruistRelay,
		RequestID:                relay.RequestID,
		IsUserRelay:              relay.IsUserRelay,
		PoktTxID:                 relay.PoktTxID.String,
		GigastakeAppID:           relay.GigastakeAppID.String,
		CreatedAt:                relay.CreatedAt,
		UpdatedAt:                relay.UpdatedAt,
		Session: types.PocketSession{
			SessionKey:    relay.SessionKey,
			SessionHeight: int(relay.SessionHeight),
			CreatedAt:     relay.CreatedAt_2,
			UpdatedAt:     relay.UpdatedAt_2,
		},
		Region: types.PortalRegion{
			PortalRegionName: relay.PortalRegionName,
		},
	}, nil
}
