package types

import "time"

type Relay struct {
	RelayID                  int64         `json:"relayID"`
	ChainID                  int32         `json:"chainID"`
	EndpointID               int32         `json:"endpointID"`
	PocketSessionID          int32         `json:"pocketSessionID"`
	PoktNodeAddress          string        `json:"poktNodeAddress"`
	RelayStartDatetime       time.Time     `json:"relayStartDatetime"`
	RelayReturnDatetime      time.Time     `json:"relayReturnDatetime"`
	IsError                  bool          `json:"isError"`
	ErrorID                  int32         `json:"errorID"`
	RelayRoundtripTime       int32         `json:"relayRoundtripTime"`
	RelayChainMethodID       int32         `json:"relayChainMethodID"`
	RelayDataSize            int32         `json:"relayDataSize"`
	RelayPortalTripTime      int32         `json:"relayPortalTripTime"`
	RelayNodeTripTime        int32         `json:"relayNodeTripTime"`
	RelayURLIsPublicEndpoint bool          `json:"relayUrlIsPublicEndpoint"`
	PortalOriginRegionID     int32         `json:"portalOriginRegionID"`
	IsAltruistRelay          bool          `json:"isAltruistRelay"`
	Error                    Error         `json:"error"`
	Session                  PocketSession `json:"session"`
	Region                   PortalRegion  `json:"region"`
	CreatedAt                time.Time     `json:"createdAt"`
	UpdatedAt                time.Time     `json:"updatedAt"`
}
