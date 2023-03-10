package types

import "time"

type ErrorType string

const (
	ErrorTypeSyncCheck  ErrorType = "sync_check"
	ErrorTypeChainCheck ErrorType = "chain_check"
	ErrorTypeRelay      ErrorType = "relay"
)

type Error struct {
	ErrorID          int64     `json:"errorID"`
	ErrorCode        int32     `json:"errorCode"`
	ErrorName        string    `json:"errorName"`
	ErrorDescription string    `json:"errorDescription"`
	ErrorType        ErrorType `json:"errorType"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
