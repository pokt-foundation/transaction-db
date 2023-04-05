CREATE TYPE error_types_enum AS ENUM ('sync_check', 'chain_check', 'relay');

CREATE  TABLE pocket_session (
	pocket_session_id    bigint  NOT NULL  GENERATED ALWAYS AS IDENTITY  ,
	session_key          varchar  NOT NULL  UNIQUE,
	session_height       integer  NOT NULL  ,
	protocol_application_id integer  NOT NULL  ,
	created_at			 date     NOT NULL  ,
	updated_at			 date     NOT NULL  ,
	CONSTRAINT pk_tbl_0 PRIMARY KEY ( pocket_session_id )
 );

CREATE  TABLE portal_region (
	portal_region_id     integer  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
	portal_region_name   varchar  NOT NULL  ,
	created_at			 date     NOT NULL  ,
	updated_at			 date     NOT NULL  ,
	CONSTRAINT pk_portal_region PRIMARY KEY ( portal_region_id )
 );

CREATE  TABLE relay (
	relay_id             bigint  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
	chain_id             integer  NOT NULL  ,
	endpoint_id          integer  NOT NULL  ,
	session_key   		varchar  NOT NULL  ,
	relay_source_url 	varchar  NOT NULL  ,
	pokt_node_address    varchar  NOT NULL  ,
	pokt_node_domain 	varchar   NOT NULL ,
	pokt_node_public_key varchar   NOT NULL ,
	relay_start_datetime date  NOT NULL  ,
	relay_return_datetime date  NOT NULL  ,
	is_error             boolean  NOT NULL  ,
	error_code           integer,
	error_name           varchar,
	error_message    	 varchar,
	error_source		 varchar,
	error_type 			 error_types_enum,
	relay_roundtrip_time integer  NOT NULL  ,
	relay_chain_method_ids varchar  NOT NULL  ,
	relay_data_size      integer  NOT NULL  ,
	relay_portal_trip_time integer  NOT NULL  ,
	relay_node_trip_time integer  NOT NULL  ,
	relay_url_is_public_endpoint boolean  NOT NULL  ,
	portal_origin_region_id integer  NOT NULL  ,
	is_altruist_relay    boolean  NOT NULL  ,
	is_user_relay 		boolean NOT NULL ,
	request_id			varchar   NOT NULL  ,
	created_at			 date     NOT NULL  ,
	updated_at			 date     NOT NULL  ,
	CONSTRAINT pk_relay PRIMARY KEY ( relay_id )
 );


ALTER TABLE relay ADD CONSTRAINT fk_relay_portal_region FOREIGN KEY ( portal_origin_region_id ) REFERENCES portal_region( portal_region_id );

ALTER TABLE relay ADD CONSTRAINT fk_relay_session FOREIGN KEY ( session_key ) REFERENCES pocket_session( session_key );
