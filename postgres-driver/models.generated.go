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

type ErrorSourcesEnum string

const (
	ErrorSourcesEnumInternal ErrorSourcesEnum = "internal"
	ErrorSourcesEnumExternal ErrorSourcesEnum = "external"
)

func (e *ErrorSourcesEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ErrorSourcesEnum(s)
	case string:
		*e = ErrorSourcesEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for ErrorSourcesEnum: %T", src)
	}
	return nil
}

type NullErrorSourcesEnum struct {
	ErrorSourcesEnum ErrorSourcesEnum
	Valid            bool // Valid is true if ErrorSourcesEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullErrorSourcesEnum) Scan(value interface{}) error {
	if value == nil {
		ns.ErrorSourcesEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ErrorSourcesEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullErrorSourcesEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ErrorSourcesEnum), nil
}

type PocketSession struct {
	ID                    int64     `json:"id"`
	SessionKey            string    `json:"sessionKey"`
	SessionHeight         int32     `json:"sessionHeight"`
	ProtocolApplicationID string    `json:"protocolApplicationID"`
	ProtocolPublicKey     string    `json:"protocolPublicKey"`
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
	ID                       int64                `json:"id"`
	RelayID                  string               `json:"relayID"`
	PoktChainID              string               `json:"poktChainID"`
	EndpointID               string               `json:"endpointID"`
	SessionKey               string               `json:"sessionKey"`
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
	PortalOriginRegionID     int32                `json:"portalOriginRegionID"`
	IsAltruistRelay          bool                 `json:"isAltruistRelay"`
	IsUserRelay              bool                 `json:"isUserRelay"`
	RequestID                string               `json:"requestID"`
	PoktTxID                 string               `json:"poktTxID"`
	CreatedAt                time.Time            `json:"createdAt"`
	UpdatedAt                time.Time            `json:"updatedAt"`
}
