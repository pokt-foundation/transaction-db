// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: query.sql

package postgresdriver

import (
	"context"
	"database/sql"
	"time"
)

const insertPocketSession = `-- name: InsertPocketSession :exec
INSERT INTO pocket_session (session_key, session_height, created_at, updated_at)
VALUES ($1, $2, $3, $4)
`

type InsertPocketSessionParams struct {
	SessionKey    string    `json:"sessionKey"`
	SessionHeight int32     `json:"sessionHeight"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func (q *Queries) InsertPocketSession(ctx context.Context, arg InsertPocketSessionParams) error {
	_, err := q.db.ExecContext(ctx, insertPocketSession,
		arg.SessionKey,
		arg.SessionHeight,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const insertPortalRegion = `-- name: InsertPortalRegion :exec
INSERT INTO portal_region (portal_region_name, created_at, updated_at)
VALUES ($1, $2, $3)
`

type InsertPortalRegionParams struct {
	PortalRegionName string    `json:"portalRegionName"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

func (q *Queries) InsertPortalRegion(ctx context.Context, arg InsertPortalRegionParams) error {
	_, err := q.db.ExecContext(ctx, insertPortalRegion, arg.PortalRegionName, arg.CreatedAt, arg.UpdatedAt)
	return err
}

const insertRelay = `-- name: InsertRelay :exec
INSERT INTO relay (
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
  created_at,
  updated_at
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29
)
`

type InsertRelayParams struct {
	PoktChainID              string               `json:"poktChainID"`
	EndpointID               string               `json:"endpointID"`
	SessionKey               string               `json:"sessionKey"`
	ProtocolAppPublicKey     string               `json:"protocolAppPublicKey"`
	RelaySourceUrl           string               `json:"relaySourceUrl"`
	PoktNodeAddress          string               `json:"poktNodeAddress"`
	PoktNodeDomain           string               `json:"poktNodeDomain"`
	PoktNodePublicKey        string               `json:"poktNodePublicKey"`
	RelayStartDatetime       time.Time            `json:"relayStartDatetime"`
	RelayReturnDatetime      time.Time            `json:"relayReturnDatetime"`
	IsError                  bool                 `json:"isError"`
	ErrorCode                sql.NullInt32        `json:"errorCode"`
	ErrorName                sql.NullString       `json:"errorName"`
	ErrorMessage             sql.NullString       `json:"errorMessage"`
	ErrorSource              NullErrorSourcesEnum `json:"errorSource"`
	ErrorType                sql.NullString       `json:"errorType"`
	RelayRoundtripTime       int32                `json:"relayRoundtripTime"`
	RelayChainMethodIds      string               `json:"relayChainMethodIds"`
	RelayDataSize            int32                `json:"relayDataSize"`
	RelayPortalTripTime      int32                `json:"relayPortalTripTime"`
	RelayNodeTripTime        int32                `json:"relayNodeTripTime"`
	RelayUrlIsPublicEndpoint bool                 `json:"relayUrlIsPublicEndpoint"`
	PortalRegionName         string               `json:"portalRegionName"`
	IsAltruistRelay          bool                 `json:"isAltruistRelay"`
	IsUserRelay              bool                 `json:"isUserRelay"`
	RequestID                string               `json:"requestID"`
	PoktTxID                 string               `json:"poktTxID"`
	CreatedAt                time.Time            `json:"createdAt"`
	UpdatedAt                time.Time            `json:"updatedAt"`
}

func (q *Queries) InsertRelay(ctx context.Context, arg InsertRelayParams) error {
	_, err := q.db.ExecContext(ctx, insertRelay,
		arg.PoktChainID,
		arg.EndpointID,
		arg.SessionKey,
		arg.ProtocolAppPublicKey,
		arg.RelaySourceUrl,
		arg.PoktNodeAddress,
		arg.PoktNodeDomain,
		arg.PoktNodePublicKey,
		arg.RelayStartDatetime,
		arg.RelayReturnDatetime,
		arg.IsError,
		arg.ErrorCode,
		arg.ErrorName,
		arg.ErrorMessage,
		arg.ErrorSource,
		arg.ErrorType,
		arg.RelayRoundtripTime,
		arg.RelayChainMethodIds,
		arg.RelayDataSize,
		arg.RelayPortalTripTime,
		arg.RelayNodeTripTime,
		arg.RelayUrlIsPublicEndpoint,
		arg.PortalRegionName,
		arg.IsAltruistRelay,
		arg.IsUserRelay,
		arg.RequestID,
		arg.PoktTxID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const selectRelay = `-- name: SelectRelay :one
SELECT r.id, r.pokt_chain_id, r.endpoint_id, r.session_key, r.protocol_app_public_key, r.relay_source_url, r.pokt_node_address, r.pokt_node_domain, r.pokt_node_public_key, r.relay_start_datetime, r.relay_return_datetime, r.is_error, r.error_code, r.error_name, r.error_message, r.error_source, r.error_type, r.relay_roundtrip_time, r.relay_chain_method_ids, r.relay_data_size, r.relay_portal_trip_time, r.relay_node_trip_time, r.relay_url_is_public_endpoint, r.is_altruist_relay, r.is_user_relay, r.request_id, r.pokt_tx_id, r.created_at, r.updated_at, ps.session_key, ps.session_height, ps.created_at, ps.updated_at, pr.portal_region_name, pr.created_at, pr.updated_at
FROM relay r
	INNER JOIN pocket_session ps ON ps.session_key = r.session_key
	INNER JOIN portal_region pr ON pr.portal_region_name = r.portal_region_name
WHERE r.id = $1
`

type SelectRelayRow struct {
	ID                       int64                `json:"id"`
	PoktChainID              string               `json:"poktChainID"`
	EndpointID               string               `json:"endpointID"`
	SessionKey               string               `json:"sessionKey"`
	ProtocolAppPublicKey     string               `json:"protocolAppPublicKey"`
	RelaySourceUrl           string               `json:"relaySourceUrl"`
	PoktNodeAddress          string               `json:"poktNodeAddress"`
	PoktNodeDomain           string               `json:"poktNodeDomain"`
	PoktNodePublicKey        string               `json:"poktNodePublicKey"`
	RelayStartDatetime       time.Time            `json:"relayStartDatetime"`
	RelayReturnDatetime      time.Time            `json:"relayReturnDatetime"`
	IsError                  bool                 `json:"isError"`
	ErrorCode                sql.NullInt32        `json:"errorCode"`
	ErrorName                sql.NullString       `json:"errorName"`
	ErrorMessage             sql.NullString       `json:"errorMessage"`
	ErrorSource              NullErrorSourcesEnum `json:"errorSource"`
	ErrorType                sql.NullString       `json:"errorType"`
	RelayRoundtripTime       int32                `json:"relayRoundtripTime"`
	RelayChainMethodIds      string               `json:"relayChainMethodIds"`
	RelayDataSize            int32                `json:"relayDataSize"`
	RelayPortalTripTime      int32                `json:"relayPortalTripTime"`
	RelayNodeTripTime        int32                `json:"relayNodeTripTime"`
	RelayUrlIsPublicEndpoint bool                 `json:"relayUrlIsPublicEndpoint"`
	IsAltruistRelay          bool                 `json:"isAltruistRelay"`
	IsUserRelay              bool                 `json:"isUserRelay"`
	RequestID                string               `json:"requestID"`
	PoktTxID                 string               `json:"poktTxID"`
	CreatedAt                time.Time            `json:"createdAt"`
	UpdatedAt                time.Time            `json:"updatedAt"`
	SessionKey_2             string               `json:"sessionKey2"`
	SessionHeight            int32                `json:"sessionHeight"`
	CreatedAt_2              time.Time            `json:"createdAt2"`
	UpdatedAt_2              time.Time            `json:"updatedAt2"`
	PortalRegionName         string               `json:"portalRegionName"`
	CreatedAt_3              time.Time            `json:"createdAt3"`
	UpdatedAt_3              time.Time            `json:"updatedAt3"`
}

func (q *Queries) SelectRelay(ctx context.Context, id int64) (SelectRelayRow, error) {
	row := q.db.QueryRowContext(ctx, selectRelay, id)
	var i SelectRelayRow
	err := row.Scan(
		&i.ID,
		&i.PoktChainID,
		&i.EndpointID,
		&i.SessionKey,
		&i.ProtocolAppPublicKey,
		&i.RelaySourceUrl,
		&i.PoktNodeAddress,
		&i.PoktNodeDomain,
		&i.PoktNodePublicKey,
		&i.RelayStartDatetime,
		&i.RelayReturnDatetime,
		&i.IsError,
		&i.ErrorCode,
		&i.ErrorName,
		&i.ErrorMessage,
		&i.ErrorSource,
		&i.ErrorType,
		&i.RelayRoundtripTime,
		&i.RelayChainMethodIds,
		&i.RelayDataSize,
		&i.RelayPortalTripTime,
		&i.RelayNodeTripTime,
		&i.RelayUrlIsPublicEndpoint,
		&i.IsAltruistRelay,
		&i.IsUserRelay,
		&i.RequestID,
		&i.PoktTxID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.SessionKey_2,
		&i.SessionHeight,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
		&i.PortalRegionName,
		&i.CreatedAt_3,
		&i.UpdatedAt_3,
	)
	return i, err
}
