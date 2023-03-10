package types

import "time"

type PocketSession struct {
	PocketSessionID       int64     `json:"pocketSessionID"`
	SessionKey            string    `json:"sessionKey"`
	SessionHeight         int32     `json:"sessionHeight"`
	ProtocolApplicationID int32     `json:"protocolApplicationID"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}
