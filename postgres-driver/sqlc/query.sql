-- name: InsertRelay :exec
INSERT INTO relay (chain_id, endpoint_id, session_key, pokt_node_address, relay_start_datetime, relay_return_datetime, is_error, error_code, error_name, error_message, error_type, relay_roundtrip_time, relay_chain_method_id, relay_data_size, relay_portal_trip_time, relay_node_trip_time, relay_url_is_public_endpoint, portal_origin_region_id, is_altruist_relay, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21);
-- name: InsertPocketSession :exec
INSERT INTO pocket_session (session_key, session_height, protocol_application_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5);
-- name: InsertPortalRegion :exec
INSERT INTO portal_region (portal_region_name, created_at, updated_at)
VALUES ($1, $2, $3);
-- name: SelectRelay :one
SELECT r.relay_id, r.chain_id, r.endpoint_id, r.session_key, r.pokt_node_address, r.relay_start_datetime, r.relay_return_datetime, r.is_error, r.error_code, r.error_name, r.error_message, r.error_type, r.relay_roundtrip_time, r.relay_chain_method_id, r.relay_data_size, r.relay_portal_trip_time, r.relay_node_trip_time, r.relay_url_is_public_endpoint, r.portal_origin_region_id, r.is_altruist_relay, r.created_at, r.updated_at, ps.session_key, ps.session_height, ps.protocol_application_id, ps.created_at, ps.updated_at, pr.portal_region_name, pr.created_at, pr.updated_at
FROM relay r
	INNER JOIN pocket_session ps ON ps.session_key = r.session_key
	INNER JOIN portal_region pr ON pr.portal_region_id = r.portal_origin_region_id
WHERE r.relay_id = $1;
