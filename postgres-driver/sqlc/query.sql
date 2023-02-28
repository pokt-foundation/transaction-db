-- name: InsertRelay :exec
INSERT INTO relay (chain_id, endpoint_id, pocket_session_id, pokt_node_address, relay_start_datetime, relay_return_datetime, is_error, error_id, relay_roundtrip_time, relay_chain_method_id, relay_data_size, relay_portal_trip_time, relay_node_trip_time, relay_url_is_public_endpoint, portal_origin_region_id, is_altruist_relay)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);
-- name: InsertError :exec
INSERT INTO error (error_code, error_name, error_description)
VALUES ($1, $2, $3);
-- name: InsertPocketSession :exec
INSERT INTO pocket_session (session_key, session_height, protocol_application_id)
VALUES ($1, $2, $3);
-- name: InsertPortalRegion :exec
INSERT INTO portal_region (portal_region_name)
VALUES ($1);
-- name: SelectRelay :one
SELECT r.relay_id, r.chain_id, r.endpoint_id, r.pocket_session_id, r.pokt_node_address, r.relay_start_datetime, r.relay_return_datetime, r.is_error, r.error_id, r.relay_roundtrip_time, r.relay_chain_method_id, r.relay_data_size, r.relay_portal_trip_time, r.relay_node_trip_time, r.relay_url_is_public_endpoint, r.portal_origin_region_id, r.is_altruist_relay, ps.session_key, ps.session_height, ps.protocol_application_id, pr.portal_region_name, e.error_code, e.error_name, e.error_description
FROM relay r
	INNER JOIN pocket_session ps ON ps.pocket_session_id = r.pocket_session_id
	INNER JOIN portal_region pr ON pr.portal_region_id = r.portal_origin_region_id
	INNER JOIN error e ON e.error_id = r.error_id
WHERE r.relay_id = $1;
