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
	ID               int64     `json:"id"`
	SessionKey       string    `json:"sessionKey"`
	SessionHeight    int32     `json:"sessionHeight"`
	PortalRegionName string    `json:"portalRegionName"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

type PortalRegion struct {
	PortalRegionName string `json:"portalRegionName"`
}

type Relay struct {
	ID                       int64                `json:"id"`
	PoktChainID              string               `json:"poktChainID"`
	EndpointID               string               `json:"endpointID"`
	SessionKey               string               `json:"sessionKey"`
	ProtocolAppPublicKey     string               `json:"protocolAppPublicKey"`
	RelaySourceUrl           string               `json:"relaySourceUrl"`
	PoktNodeAddress          sql.NullString       `json:"poktNodeAddress"`
	PoktNodeDomain           sql.NullString       `json:"poktNodeDomain"`
	PoktNodePublicKey        sql.NullString       `json:"poktNodePublicKey"`
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
	PoktTxID                 sql.NullString       `json:"poktTxID"`
	CreatedAt                time.Time            `json:"createdAt"`
	UpdatedAt                time.Time            `json:"updatedAt"`
}

type ServiceRecord struct {
	ID                     int64     `json:"id"`
	NodePublicKey          string    `json:"nodePublicKey"`
	PoktChainID            string    `json:"poktChainID"`
	SessionKey             string    `json:"sessionKey"`
	RequestID              string    `json:"requestID"`
	PortalRegionName       string    `json:"portalRegionName"`
	Latency                float64   `json:"latency"`
	Tickets                int32     `json:"tickets"`
	Result                 string    `json:"result"`
	Available              bool      `json:"available"`
	Successes              int32     `json:"successes"`
	Failures               int32     `json:"failures"`
	P90SuccessLatency      float64   `json:"p90SuccessLatency"`
	MedianSuccessLatency   float64   `json:"medianSuccessLatency"`
	WeightedSuccessLatency float64   `json:"weightedSuccessLatency"`
	SuccessRate            float64   `json:"successRate"`
	CreatedAt              time.Time `json:"createdAt"`
	UpdatedAt              time.Time `json:"updatedAt"`
}
