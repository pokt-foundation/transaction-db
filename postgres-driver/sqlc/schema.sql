CREATE  TABLE error (
	error_id             integer  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
	error_code           integer  NOT NULL  ,
	error_name           varchar  NOT NULL  ,
	error_description    varchar  NOT NULL  ,
	CONSTRAINT pk_error PRIMARY KEY ( error_id )
 );

CREATE  TABLE pocket_session (
	pocket_session_id    integer  NOT NULL  GENERATED ALWAYS AS IDENTITY  ,
	session_key          varchar  NOT NULL  ,
	session_height       integer  NOT NULL  ,
	protocol_application_id integer  NOT NULL  ,
	CONSTRAINT pk_tbl_0 PRIMARY KEY ( pocket_session_id )
 );

CREATE  TABLE portal_region (
	portal_region_id     integer  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
	portal_region_name   varchar  NOT NULL  ,
	CONSTRAINT pk_portal_region PRIMARY KEY ( portal_region_id )
 );

CREATE  TABLE relay (
	relay_id             bigint  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
	chain_id             integer  NOT NULL  ,
	endpoint_id          integer  NOT NULL  ,
	pocket_session_id    integer  NOT NULL  ,
	pokt_node_address    varchar  NOT NULL  ,
	relay_start_datetime date  NOT NULL  ,
	relay_return_datetime date  NOT NULL  ,
	is_error             boolean  NOT NULL  ,
	error_id             integer  NOT NULL  ,
	relay_roundtrip_time integer  NOT NULL  ,
	relay_chain_method_id integer  NOT NULL  ,
	relay_data_size      integer  NOT NULL  ,
	relay_portal_trip_time integer  NOT NULL  ,
	relay_node_trip_time integer  NOT NULL  ,
	relay_url_is_public_endpoint boolean  NOT NULL  ,
	portal_origin_region_id integer  NOT NULL  ,
	is_altruist_relay    boolean  NOT NULL  ,
	CONSTRAINT pk_relay PRIMARY KEY ( relay_id )
 );


ALTER TABLE relay ADD CONSTRAINT fk_relay_portal_region FOREIGN KEY ( portal_origin_region_id ) REFERENCES portal_region( portal_region_id );

ALTER TABLE relay ADD CONSTRAINT fk_relay_error FOREIGN KEY ( error_id ) REFERENCES error( error_id );

ALTER TABLE relay ADD CONSTRAINT fk_relay_session FOREIGN KEY ( pocket_session_id ) REFERENCES pocket_session( pocket_session_id );
