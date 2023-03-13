package types

import "time"

type ErrorType string

const (
	ErrorTypeSyncCheck  ErrorType = "sync_check"
	ErrorTypeChainCheck ErrorType = "chain_check"
	ErrorTypeRelay      ErrorType = "relay"
)

type Relay struct {
	RelayID                  int64         `json:"relayID"`
	ChainID                  int32         `json:"chainID"`
	EndpointID               int32         `json:"endpointID"`
	SessionKey               string        `json:"sessionKey"`
	PoktNodeAddress          string        `json:"poktNodeAddress"`
	RelayStartDatetime       time.Time     `json:"relayStartDatetime"`
	RelayReturnDatetime      time.Time     `json:"relayReturnDatetime"`
	IsError                  bool          `json:"isError"`
	ErrorCode                int32         `json:"errorCode,omitempty"`
	ErrorName                string        `json:"errorName,omitempty"`
	ErrorMessage             string        `json:"errorMessage,omitempty"`
	ErrorType                ErrorType     `json:"errorType,omitempty"`
	RelayRoundtripTime       int32         `json:"relayRoundtripTime"`
	RelayChainMethodID       int32         `json:"relayChainMethodID"`
	RelayDataSize            int32         `json:"relayDataSize"`
	RelayPortalTripTime      int32         `json:"relayPortalTripTime"`
	RelayNodeTripTime        int32         `json:"relayNodeTripTime"`
	RelayURLIsPublicEndpoint bool          `json:"relayUrlIsPublicEndpoint"`
	PortalOriginRegionID     int32         `json:"portalOriginRegionID"`
	IsAltruistRelay          bool          `json:"isAltruistRelay"`
	Session                  PocketSession `json:"session"`
	Region                   PortalRegion  `json:"region"`
	CreatedAt                time.Time     `json:"createdAt"`
	UpdatedAt                time.Time     `json:"updatedAt"`
}
