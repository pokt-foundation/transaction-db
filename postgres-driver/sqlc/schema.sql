CREATE TYPE error_sources_enum AS ENUM ('internal', 'external');

CREATE  TABLE pocket_session (
	id    				 bigint  NOT NULL  GENERATED ALWAYS AS IDENTITY  ,
	session_key          char(44)  NOT NULL  UNIQUE,
	session_height       integer  NOT NULL  ,
	created_at			 date     NOT NULL  ,
	updated_at			 date     NOT NULL  ,
	CONSTRAINT pk_tbl_0 PRIMARY KEY ( id )
 );

CREATE  TABLE portal_region (
	portal_region_name   varchar  NOT NULL  ,
	created_at			 date     NOT NULL  ,
	updated_at			 date     NOT NULL  ,
	CONSTRAINT pk_portal_region PRIMARY KEY ( portal_region_name )
 );

CREATE  TABLE relay (
	id             		bigint  NOT NULL GENERATED ALWAYS AS IDENTITY  ,
	pokt_chain_id       varchar  NOT NULL  ,
	endpoint_id          varchar  NOT NULL  ,
	session_key   		char(44)  NOT NULL  ,
	protocol_app_public_key	 char(64) NOT NULL ,
	relay_source_url 	varchar  NOT NULL  ,
	pokt_node_address    char(40)  NOT NULL  ,
	pokt_node_domain 	varchar   NOT NULL ,
	pokt_node_public_key char(64)   NOT NULL ,
	relay_start_datetime date  NOT NULL  ,
	relay_return_datetime date  NOT NULL  ,
	is_error             boolean  NOT NULL  ,
	error_code           integer,
	error_name           varchar,
	error_message    	 varchar,
	error_source		 error_sources_enum,
	error_type 			 varchar,
	relay_roundtrip_time integer  NOT NULL  ,
	relay_chain_method_ids varchar  NOT NULL  ,
	relay_data_size      integer  NOT NULL  ,
	relay_portal_trip_time integer  NOT NULL  ,
	relay_node_trip_time integer  NOT NULL  ,
	relay_url_is_public_endpoint boolean  NOT NULL  ,
	portal_region_name varchar  NOT NULL  ,
	is_altruist_relay    boolean  NOT NULL  ,
	is_user_relay 		boolean NOT NULL ,
	request_id			varchar   NOT NULL  ,
	pokt_tx_id 			varchar	  NOT NULL  ,
	created_at			 date     NOT NULL  ,
	updated_at			 date     NOT NULL  ,
	CONSTRAINT pk_relay PRIMARY KEY ( id )
 );


ALTER TABLE relay ADD CONSTRAINT fk_relay_portal_region FOREIGN KEY ( portal_region_name ) REFERENCES portal_region( portal_region_name );

ALTER TABLE relay ADD CONSTRAINT fk_relay_session FOREIGN KEY ( session_key ) REFERENCES pocket_session( session_key );
