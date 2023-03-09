package types

import "time"

type Error struct {
	ErrorID          int64     `json:"errorID"`
	ErrorCode        int32     `json:"errorCode"`
	ErrorName        string    `json:"errorName"`
	ErrorDescription string    `json:"errorDescription"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
