// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package postgresdriver

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type ErrorTypesEnum string

const (
	ErrorTypesEnumSyncCheck  ErrorTypesEnum = "sync_check"
	ErrorTypesEnumChainCheck ErrorTypesEnum = "chain_check"
	ErrorTypesEnumRelay      ErrorTypesEnum = "relay"
)

func (e *ErrorTypesEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ErrorTypesEnum(s)
	case string:
		*e = ErrorTypesEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for ErrorTypesEnum: %T", src)
	}
	return nil
}

type NullErrorTypesEnum struct {
	ErrorTypesEnum ErrorTypesEnum
	Valid          bool // Valid is true if ErrorTypesEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullErrorTypesEnum) Scan(value interface{}) error {
	if value == nil {
		ns.ErrorTypesEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ErrorTypesEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullErrorTypesEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ErrorTypesEnum), nil
}

type PocketSession struct {
	PocketSessionID       int64     `json:"pocketSessionID"`
	SessionKey            string    `json:"sessionKey"`
	SessionHeight         int32     `json:"sessionHeight"`
	ProtocolApplicationID int32     `json:"protocolApplicationID"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}

type PortalRegion struct {
	PortalRegionID   int32     `json:"portalRegionID"`
	PortalRegionName string    `json:"portalRegionName"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type Relay struct {
	RelayID                  int64              `json:"relayID"`
	ChainID                  int32              `json:"chainID"`
	EndpointID               int32              `json:"endpointID"`
	SessionKey               string             `json:"sessionKey"`
	PoktNodeAddress          string             `json:"poktNodeAddress"`
	RelayStartDatetime       time.Time          `json:"relayStartDatetime"`
	RelayReturnDatetime      time.Time          `json:"relayReturnDatetime"`
	IsError                  bool               `json:"isError"`
	ErrorCode                sql.NullInt32      `json:"errorCode"`
	ErrorName                sql.NullString     `json:"errorName"`
	ErrorMessage             sql.NullString     `json:"errorMessage"`
	ErrorType                NullErrorTypesEnum `json:"errorType"`
	RelayRoundtripTime       int32              `json:"relayRoundtripTime"`
	RelayChainMethodID       int32              `json:"relayChainMethodID"`
	RelayDataSize            int32              `json:"relayDataSize"`
	RelayPortalTripTime      int32              `json:"relayPortalTripTime"`
	RelayNodeTripTime        int32              `json:"relayNodeTripTime"`
	RelayUrlIsPublicEndpoint bool               `json:"relayUrlIsPublicEndpoint"`
	PortalOriginRegionID     int32              `json:"portalOriginRegionID"`
	IsAltruistRelay          bool               `json:"isAltruistRelay"`
	CreatedAt                time.Time          `json:"createdAt"`
	UpdatedAt                time.Time          `json:"updatedAt"`
}